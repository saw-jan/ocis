package cs3

import (
	"context"
	"crypto/tls"
	"fmt"

	cs3gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	cs3rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	"github.com/libregraph/lico"
	"github.com/libregraph/lico/config"
	"github.com/libregraph/lico/identifier/backends"
	"github.com/libregraph/lico/identifier/meta/scopes"
	"github.com/libregraph/lico/identity"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	ins "google.golang.org/grpc/credentials/insecure"
	"stash.kopano.io/kgol/oidc-go"
)

const cs3BackendName = "identifier-cs3"

var cs3SpportedScopes = []string{
	oidc.ScopeProfile,
	oidc.ScopeEmail,
	lico.ScopeUniqueUserID,
	lico.ScopeRawSubject,
}

// CS3 Backend holds the data for the CS3 identifier backend
type CS3Backend struct {
	supportedScopes []string

	logger            logrus.FieldLogger
	tlsConfig         *tls.Config
	gatewayURI        string
	machineAuthAPIKey string
	insecure          bool

	sessions cmap.ConcurrentMap

	gateway cs3gateway.GatewayAPIClient
}

// NewCS3Backend creates a new CS3 backend identifier backend
func NewCS3Backend(
	c *config.Config,
	tlsConfig *tls.Config,
	gatewayURI string,
	machineAuthAPIKey string,
	insecure bool,
) (*CS3Backend, error) {

	// Build supported scopes based on default scopes.
	supportedScopes := make([]string, len(cs3SpportedScopes))
	copy(supportedScopes, cs3SpportedScopes)

	b := &CS3Backend{
		supportedScopes: supportedScopes,

		logger:            c.Logger,
		tlsConfig:         tlsConfig,
		gatewayURI:        gatewayURI,
		machineAuthAPIKey: machineAuthAPIKey,
		insecure:          insecure,

		sessions: cmap.New(),
	}

	b.logger.Infoln("cs3 backend connection set up")

	return b, nil
}

// RunWithContext implements the Backend interface.
func (b *CS3Backend) RunWithContext(ctx context.Context) error {
	return nil
}

// Logon implements the Backend interface, enabling Logon with user name and
// password as provided. Requests are bound to the provided context.
func (b *CS3Backend) Logon(ctx context.Context, audience, username, password string) (bool, *string, *string, backends.UserFromBackend, error) {

	l, err := b.connect(ctx)
	if err != nil {
		return false, nil, nil, nil, fmt.Errorf("cs3 backend logon connect error: %v", err)
	}
	defer l.Close()

	client := cs3gateway.NewGatewayAPIClient(l)

	res, err := client.Authenticate(ctx, &cs3gateway.AuthenticateRequest{
		Type:         "basic",
		ClientId:     username,
		ClientSecret: password,
	})
	if err != nil {
		return false, nil, nil, nil, fmt.Errorf("cs3 backend basic authenticate rpc error: %v", err)
	}
	if res.Status.Code != cs3rpc.Code_CODE_OK {
		return false, nil, nil, nil, fmt.Errorf("cs3 backend basic authenticate failed with code %s: %s", res.Status.Code.String(), res.Status.Message)
	}

	session := createSession(ctx, res.User)

	user, err := newCS3User(res.User)
	if err != nil {
		return false, nil, nil, nil, fmt.Errorf("cs3 backend resolve entry data error: %v", err)
	}

	// Use the users subject as user id.
	userID := user.Subject()

	sessionRef := identity.GetSessionRef(b.Name(), audience, userID)
	b.sessions.Set(*sessionRef, session)
	b.logger.WithFields(logrus.Fields{
		"session":  session,
		"ref":      *sessionRef,
		"username": user.Username(),
		"id":       userID,
	}).Debugln("cs3 backend logon")

	return true, &userID, sessionRef, user, nil
}

// GetUser implements the Backend interface, providing user meta data retrieval
// for the user specified by the userID. Requests are bound to the provided
// context.
func (b *CS3Backend) GetUser(ctx context.Context, userEntryID string, sessionRef *string, requestedScopes map[string]bool) (backends.UserFromBackend, error) {

	session, err := b.getSessionForUser(ctx, userEntryID, sessionRef, true, true, false)
	if err != nil {
		return nil, fmt.Errorf("cs3 backend resolve session error: %v", err)
	}

	user, err := newCS3User(session.User())
	if err != nil {
		return nil, fmt.Errorf("cs3 backend get user failed to process user: %v", err)
	}
	// TODO double check userEntryID matches session?

	return user, nil
}

// ResolveUserByUsername implements the Backend interface, providing lookup for
// user by providing the username. Requests are bound to the provided context.
func (b *CS3Backend) ResolveUserByUsername(ctx context.Context, username string) (backends.UserFromBackend, error) {

	l, err := b.connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("cs3 backend resolve username connect error: %v", err)
	}
	defer l.Close()

	client := cs3gateway.NewGatewayAPIClient(l)

	res, err := client.Authenticate(ctx, &cs3gateway.AuthenticateRequest{
		Type:         "machine",
		ClientId:     "username:" + username,
		ClientSecret: b.machineAuthAPIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("cs3 backend machine authenticate rpc error: %v", err)
	}
	if res.Status.Code != cs3rpc.Code_CODE_OK {
		return nil, fmt.Errorf("cs3 backend machine authenticate failed with code %s: %s", res.Status.Code.String(), res.Status.Message)
	}

	user, err := newCS3User(res.User)
	if err != nil {
		return nil, fmt.Errorf("cs3 backend resolve username data error: %v", err)
	}

	return user, nil
}

// RefreshSession implements the Backend interface.
func (b *CS3Backend) RefreshSession(ctx context.Context, userID string, sessionRef *string, claims map[string]interface{}) error {
	return nil
}

// DestroySession implements the Backend interface providing destroy CS3 session.
func (b *CS3Backend) DestroySession(ctx context.Context, sessionRef *string) error {
	b.sessions.Remove(*sessionRef)
	return nil
}

// UserClaims implements the Backend interface, providing user specific claims
// for the user specified by the userID.
func (b *CS3Backend) UserClaims(userID string, authorizedScopes map[string]bool) map[string]interface{} {
	return nil
	// TODO should we return the "ownclouduuid" as a claim? there is also "LibgreGraph.UUID" / lico.ScopeUniqueUserID
}

// ScopesSupported implements the Backend interface, providing supported scopes
// when running this backend.
func (b *CS3Backend) ScopesSupported() []string {
	return b.supportedScopes
}

// ScopesMeta implements the Backend interface, providing meta data for
// supported scopes.
func (b *CS3Backend) ScopesMeta() *scopes.Scopes {
	return nil
}

// Name implements the Backend interface.
func (b *CS3Backend) Name() string {
	return cs3BackendName
}

func (b *CS3Backend) connect(ctx context.Context) (*grpc.ClientConn, error) {
	if b.insecure {
		return grpc.Dial(b.gatewayURI, grpc.WithTransportCredentials(ins.NewCredentials()))
	}

	creds := credentials.NewTLS(b.tlsConfig)
	return grpc.Dial(b.gatewayURI, grpc.WithTransportCredentials(creds))
}

func (b *CS3Backend) getSessionForUser(ctx context.Context, userEntryID string, sessionRef *string, register bool, refresh bool, removeIfRegistered bool) (*cs3Session, error) {
	if sessionRef == nil {
		return nil, nil
	}

	var session *cs3Session
	if s, ok := b.sessions.Get(*sessionRef); ok {
		// Existing session.
		session = s.(*cs3Session)
		if session != nil {
			return session, nil
		}
	}

	return session, nil
}
