package identity

import (
	"context"
	"net/url"

	"github.com/CiscoM31/godata"
	cs3group "github.com/cs3org/go-cs3apis/cs3/identity/group/v1beta1"
	cs3user "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	libregraph "github.com/owncloud/libre-graph-api-go"
	"github.com/owncloud/ocis/v2/services/graph/pkg/errorcode"
)

// Errors used by the interfaces
var (
	// ErrReadOnly signals that the backend is set to read only.
	ErrReadOnly = errorcode.New(errorcode.NotAllowed, "server is configured read-only")
	// ErrNotFound signals that the requested resource was not found.
	ErrNotFound = errorcode.New(errorcode.ItemNotFound, "not found")
)

// Backend defines the Interface for an IdentityBackend implementation
type Backend interface {
	// CreateUser creates a given user in the identity backend.
	CreateUser(ctx context.Context, user libregraph.User) (*libregraph.User, error)
	// DeleteUser deletes a given user, identified by username or id, from the backend
	DeleteUser(ctx context.Context, nameOrID string) error
	// UpdateUser applies changes to given user, identified by username or id
	UpdateUser(ctx context.Context, nameOrID string, user libregraph.User) (*libregraph.User, error)
	GetUser(ctx context.Context, nameOrID string, oreq *godata.GoDataRequest) (*libregraph.User, error)
	GetUsers(ctx context.Context, oreq *godata.GoDataRequest) ([]*libregraph.User, error)

	// CreateGroup creates the supplied group in the identity backend.
	CreateGroup(ctx context.Context, group libregraph.Group) (*libregraph.Group, error)
	// DeleteGroup deletes a given group, identified by id
	DeleteGroup(ctx context.Context, id string) error
	// UpdateGroupName updates the group name
	UpdateGroupName(ctx context.Context, groupID string, groupName string) error
	GetGroup(ctx context.Context, nameOrID string, queryParam url.Values) (*libregraph.Group, error)
	GetGroups(ctx context.Context, queryParam url.Values) ([]*libregraph.Group, error)
	// GetGroupMembers list all members of a group
	GetGroupMembers(ctx context.Context, id string, oreq *godata.GoDataRequest) ([]*libregraph.User, error)
	// AddMembersToGroup adds new members (reference by a slice of IDs) to supplied group in the identity backend.
	AddMembersToGroup(ctx context.Context, groupID string, memberID []string) error
	// RemoveMemberFromGroup removes a single member (by ID) from a group
	RemoveMemberFromGroup(ctx context.Context, groupID string, memberID string) error
}

// EducationBackend defines the Interface for an EducationBackend implementation
type EducationBackend interface {
	// CreateEducationSchool creates the supplied school in the identity backend.
	CreateEducationSchool(ctx context.Context, group libregraph.EducationSchool) (*libregraph.EducationSchool, error)
	// DeleteSchool deletes a given school, identified by id
	DeleteEducationSchool(ctx context.Context, id string) error
	// GetEducationSchool reads a given school by id
	GetEducationSchool(ctx context.Context, nameOrID string) (*libregraph.EducationSchool, error)
	// GetEducationSchools lists all schools
	GetEducationSchools(ctx context.Context) ([]*libregraph.EducationSchool, error)
	// UpdateEducationSchool updates attributes of a school
	UpdateEducationSchool(ctx context.Context, numberOrID string, school libregraph.EducationSchool) (*libregraph.EducationSchool, error)
	// GetEducationSchoolUsers lists all members of a school
	GetEducationSchoolUsers(ctx context.Context, id string) ([]*libregraph.EducationUser, error)
	// AddUsersToEducationSchool adds new members (reference by a slice of IDs) to supplied school in the identity backend.
	AddUsersToEducationSchool(ctx context.Context, schoolID string, memberID []string) error
	// RemoveUserFromEducationSchool removes a single member (by ID) from a school
	RemoveUserFromEducationSchool(ctx context.Context, schoolID string, memberID string) error

	// GetEducationSchoolClasses lists all classes in a chool
	GetEducationSchoolClasses(ctx context.Context, schoolNumberOrID string) ([]*libregraph.EducationClass, error)
	// AddClassesToEducationSchool adds new classes (referenced by a slice of IDs) to supplied school in the identity backend.
	AddClassesToEducationSchool(ctx context.Context, schoolNumberOrID string, memberIDs []string) error
	// RemoveClassFromEducationSchool removes a class from a school.
	RemoveClassFromEducationSchool(ctx context.Context, schoolNumberOrID string, memberID string) error

	// GetEducationClasses lists all classes
	GetEducationClasses(ctx context.Context) ([]*libregraph.EducationClass, error)
	// GetEducationClasses reads a given class by id
	GetEducationClass(ctx context.Context, namedOrID string) (*libregraph.EducationClass, error)
	// CreateEducationClass creates the supplied education class in the identity backend.
	CreateEducationClass(ctx context.Context, class libregraph.EducationClass) (*libregraph.EducationClass, error)
	// DeleteEducationClass deletes the supplied education class in the identity backend.
	DeleteEducationClass(ctx context.Context, nameOrID string) error
	// GetEducationClassMembers returns the EducationUser members for an EducationClass
	GetEducationClassMembers(ctx context.Context, nameOrID string) ([]*libregraph.EducationUser, error)
	// UpdateEducationClass updates properties of the supplied class in the identity backend.
	UpdateEducationClass(ctx context.Context, id string, class libregraph.EducationClass) (*libregraph.EducationClass, error)

	// CreateEducationUser creates a given education user in the identity backend.
	CreateEducationUser(ctx context.Context, user libregraph.EducationUser) (*libregraph.EducationUser, error)
	// DeleteEducationUser deletes a given educationuser, identified by username or id, from the backend
	DeleteEducationUser(ctx context.Context, nameOrID string) error
	// UpdateEducationUser applies changes to given education user, identified by username or id
	UpdateEducationUser(ctx context.Context, nameOrID string, user libregraph.EducationUser) (*libregraph.EducationUser, error)
	// GetEducationUser reads an education user by id or name
	GetEducationUser(ctx context.Context, nameOrID string) (*libregraph.EducationUser, error)
	// GetEducationUsers lists all education users
	GetEducationUsers(ctx context.Context) ([]*libregraph.EducationUser, error)

	// GetEducationClassTeachers returns the EducationUser teachers for an EducationClass
	GetEducationClassTeachers(ctx context.Context, classID string) ([]*libregraph.EducationUser, error)
	// AddTeacherToEducationclass adds a teacher (by ID) to class in the identity backend.
	AddTeacherToEducationClass(ctx context.Context, classID string, teacherID string) error
	// RemoveTeacherFromEducationClass removes teacher (by ID) from a class
	RemoveTeacherFromEducationClass(ctx context.Context, classID string, teacherID string) error
}

// CreateUserModelFromCS3 converts a cs3 User object into a libregraph.User
func CreateUserModelFromCS3(u *cs3user.User) *libregraph.User {
	if u.Id == nil {
		u.Id = &cs3user.UserId{}
	}
	return &libregraph.User{
		DisplayName:              &u.DisplayName,
		Mail:                     &u.Mail,
		OnPremisesSamAccountName: &u.Username,
		Id:                       &u.Id.OpaqueId,
	}
}

// CreateGroupModelFromCS3 converts a cs3 Group object into a libregraph.Group
func CreateGroupModelFromCS3(g *cs3group.Group) *libregraph.Group {
	if g.Id == nil {
		g.Id = &cs3group.GroupId{}
	}
	return &libregraph.Group{
		Id:          &g.Id.OpaqueId,
		DisplayName: &g.GroupName,
	}
}
