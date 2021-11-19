# images
TRAEFIK_IMG = "traefik:v2.5"
OCIS_IMG = "owncloud/ocis:latest"
OC10_IMG = "owncloud/server:latest"
OC10_DB_IMG = "mariadb:10.6"
KEYCLOAK_IMG = "quay.io/keycloak/keycloak:latest"
KEYCLOAK_DB_IMG = "postgres:alpine"
OPENLDAP_IMG = "osixia/openldap:latest"
LDAP_MANAGER_IMG = "osixia/phpldapadmin:0.9.0"
REDIS_IMG = "redis:6"

def main(ctx):
    pipelines = []
    pipelines = 
    return pipelines

def parallelDeployAcceptance():
    return {
        "kind": "pipeline",
        "type": "docker",
        "name": "acceptanceTests",
        "platform": {
            "os": "linux",
            "arch": "amd64",
        },
        "steps":
        "services":
        "volumes":
    }

def traefikService():
    return [{
        "name": "traefik"
    }]

def ocis():
    environment = {
      "PROXY_ENABLE_BASIC_AUTH": "true"
      # Keycloak IDP specific configuration
      "PROXY_OIDC_ISSUER": https://${KEYCLOAK_DOMAIN:-keycloak.owncloud.test}/auth/realms/${KEYCLOAK_REALM:-owncloud}
      "WEB_OIDC_AUTHORITY": https://${KEYCLOAK_DOMAIN:-keycloak.owncloud.test}/auth/realms/${KEYCLOAK_REALM:-owncloud}
      "WEB_OIDC_CLIENT_ID": ocis-web
      "WEB_OIDC_METADATA_URL": https://${KEYCLOAK_DOMAIN:-keycloak.owncloud.test}/auth/realms/${KEYCLOAK_REALM:-owncloud}/.well-known/openid-configuration
      "STORAGE_OIDC_ISSUER": https://${KEYCLOAK_DOMAIN:-keycloak.owncloud.test}
      "STORAGE_LDAP_IDP": https://${KEYCLOAK_DOMAIN:-keycloak.owncloud.test}/auth/realms/${KEYCLOAK_REALM:-owncloud}
      "WEB_OIDC_SCOPE": openid profile email owncloud
      # LDAP bind
      "STORAGE_LDAP_HOSTNAME": openldap
      "STORAGE_LDAP_PORT": 636
      "STORAGE_LDAP_INSECURE": "true"
      "STORAGE_LDAP_BIND_DN": "cn=admin,dc=owncloud,dc=com"
      "STORAGE_LDAP_BIND_PASSWORD": ${LDAP_ADMIN_PASSWORD:-admin}
      # LDAP user settings
      "PROXY_AUTOPROVISION_ACCOUNTS": "true" # automatically create users when they login
      "PROXY_ACCOUNT_BACKEND_TYPE": cs3 # proxy should get users from CS3APIS (which gets it from LDAP)
      "PROXY_USER_OIDC_CLAIM": ocis.user.uuid # claim was added in Keycloak
      "PROXY_USER_CS3_CLAIM": userid # equals STORAGE_LDAP_USER_SCHEMA_UID
      "STORAGE_LDAP_BASE_DN": "dc=owncloud,dc=com"
      "STORAGE_LDAP_GROUP_SCHEMA_DISPLAYNAME": "cn"
      "STORAGE_LDAP_GROUP_SCHEMA_GID_NUMBER": "gidnumber"
      "STORAGE_LDAP_GROUP_SCHEMA_GID": "cn"
      "STORAGE_LDAP_GROUP_SCHEMA_MAIL": "mail"
      "STORAGE_LDAP_GROUPATTRIBUTEFILTER": "(&(objectclass=posixGroup)(objectclass=owncloud)({{attr}}={{value}}))"
      "STORAGE_LDAP_GROUPFILTER": "(&(objectclass=groupOfUniqueNames)(objectclass=owncloud)(ownclouduuid={{.OpaqueId}}*))"
      "STORAGE_LDAP_GROUPMEMBERFILTER": "(&(objectclass=posixAccount)(objectclass=owncloud)(ownclouduuid={{.OpaqueId}}*))"
      "STORAGE_LDAP_USERGROUPFILTER": "(&(objectclass=posixGroup)(objectclass=owncloud)(ownclouduuid={{.OpaqueId}}*))"
      "STORAGE_LDAP_USER_SCHEMA_CN": "cn"
      "STORAGE_LDAP_USER_SCHEMA_DISPLAYNAME": "displayname"
      "STORAGE_LDAP_USER_SCHEMA_GID_NUMBER": "gidnumber"
      "STORAGE_LDAP_USER_SCHEMA_MAIL": "mail"
      "STORAGE_LDAP_USER_SCHEMA_UID_NUMBER": "uidnumber"
      "STORAGE_LDAP_USER_SCHEMA_UID": "ownclouduuid"
      "STORAGE_LDAP_LOGINFILTER": "(&(objectclass=posixAccount)(objectclass=owncloud)(|(uid={{login}})(mail={{login}})))"
      "STORAGE_LDAP_USERATTRIBUTEFILTER": "(&(objectclass=posixAccount)(objectclass=owncloud)({{attr}}={{value}}))"
      "STORAGE_LDAP_USERFILTER": "(&(objectclass=posixAccount)(objectclass=owncloud)(|(ownclouduuid={{.OpaqueId}})(uid={{.OpaqueId}})))"
      "STORAGE_LDAP_USERFINDFILTER": "(&(objectclass=posixAccount)(objectclass=owncloud)(|(cn={{query}}*)(displayname={{query}}*)(mail={{query}}*)))"
      # ownCloud storage driver
      "STORAGE_HOME_DRIVER": owncloudsql
      "STORAGE_USERS_DRIVER": owncloudsql
      "STORAGE_METADATA_DRIVER": ocis # keep metadata on ocis storage since this are only small files atm
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DATADIR": /mnt/data/files
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_UPLOADINFO_DIR": /tmp
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_SHARE_FOLDER": "/Shares"
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_LAYOUT": "{{.Username}}"
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DBUSERNAME": owncloud
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DBPASSWORD": owncloud
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DBHOST": oc10-db
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DBPORT": 3306
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_DBNAME": owncloud
      "STORAGE_USERS_DRIVER_OWNCLOUDSQL_REDIS_ADDR": redis:6379 # TODO: redis is not yet supported
      # ownCloud storage readonly
      "OCIS_STORAGE_READ_ONLY": "false" # TODO: conflict with OWNCLOUDSQL -> https://github.com/owncloud/ocis/issues/2303
      # General oCIS config
      "OCIS_LOG_LEVEL": ${OCIS_LOG_LEVEL:-error} # make oCIS less verbose
      "PROXY_LOG_LEVEL": ${PROXY_LOG_LEVEL:-error}
      "OCIS_URL": https://${CLOUD_DOMAIN:-cloud.owncloud.test}
      "PROXY_TLS": "false" # do not use SSL between Traefik and oCIS
      "PROXY_CONFIG_FILE": "/var/tmp/ocis/.config/proxy-config.json"
      # change default secrets
      "OCIS_JWT_SECRET": ${OCIS_JWT_SECRET:-Pive-Fumkiu4}
      "STORAGE_TRANSFER_SECRET": ${STORAGE_TRANSFER_SECRET:-replace-me-with-a-transfer-secret}
      "OCIS_MACHINE_AUTH_API_KEY": ${OCIS_MACHINE_AUTH_API_KEY:-change-me-please}
      # INSECURE: needed if oCIS / Traefik is using self generated certificates
      "OCIS_INSECURE": "false"
    }

    return [{
        "name": "ocis",
        "image": OCIS_IMG,
        "pull": "always",
        "detach": True,
        "environment": environment
    }]

def redis():
    return [{
        "name": "redis",
        "image": REDIS_IMG,
        "volumes": [{
            "name": "oc10-redis-data",
            "path": "/data",
        }]
    }]