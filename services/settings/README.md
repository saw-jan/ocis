# Settings

The `settings` service provides functionality for other services to register new settings as well as storing and retrieving the respective settings' values.

## Settings Managed

The settings service is currently used for managing the:

*   users' `profile` settings like the language and the email notification settings,
*   possible user roles and their respective permissions,
*   assignment of roles to users.

As an example, user profile settings that can be changed in the Web UI must be persistent.

The settings service supports two different backends for persisting the data. The backend can be set via the `SETTINGS_STORE_TYPE` environment variable. Supported values are:

*   `metadata`: The default. This backend persists the settings data via the `storage-system` service.
*   `filesystem`: This backend persists the settings data in a directory on the local filesystem.
  The directory can be configured with `SETTINGS_DATA_PATH`. This backend is **not** suitable for running
  multiple intances of the `settings` service in a scale-out deployment and should be therefore considered
  deprecated.

<!--- Note: The diagramm is outdate, leaving it here for a future rework
The diagram shows how the settings service integrates into oCIS:

The diagram shows how the settings service integrates into oCIS:

```mermaid
graph TD
    ows ---|"listSettingsBundles(),<br>saveSettingsValue(value)"| os[ocis-settings]
    owc ---|"listSettingsValues()"| sdk[oC SDK]
    sdk --- sdks{ocis-settings<br>available?}
    sdks ---|"yes"| os
    sdks ---|"no"| defaults[Use set of<br>default values]
    oa[oCIS services<br>e.g. ocis-accounts] ---|"saveSettingsBundle(bundle)"| os
```
-->

## Caching

When using `SETTINGS_STORE_TYPE=metadata`, the `settings` service caches the results of queries against the storage backend to provide faster responses. The content of this cache is independent of the cache used in the `storage-system` service as it caches directory listing and settings content stored in files.

The store used for the cache can be configured using the `SETTINGS_CACHE_STORE` environment variable. Possible stores are:
  -   `memory`: Basic in-memory store and the default.
  -   `redis-sentinel`: Stores data in a configured Redis Sentinel cluster.
  -   `nats-js-kv`: Stores data using key-value-store feature of [nats jetstream](https://docs.nats.io/nats-concepts/jetstream/key-value-store)
  -   `noop`: Stores nothing. Useful for testing. Not recommended in production environments.
  -   `ocmem`: Advanced in-memory store allowing max size. (deprecated)
  -   `redis`: Stores data in a configured Redis cluster. (deprecated)
  -   `etcd`: Stores data in a configured etcd cluster. (deprecated)
  -   `nats-js`: Stores data using object-store feature of [nats jetstream](https://docs.nats.io/nats-concepts/jetstream/obj_store) (deprecated)

Other store types may work but are not supported currently.

Note: The service can only be scaled if not using `memory` store and the stores are configured identically over all instances!


Note that if you have used one of the deprecated stores, you should reconfigure to one of the supported ones as the deprecated stores will be removed in a later version.

Store specific notes:
  -   When using `redis-sentinel`, the Redis master to use is configured via e.g. `OCIS_CACHE_STORE_NODES` in the form of `<sentinel-host>:<sentinel-port>/<redis-master>` like `10.10.0.200:26379/mymaster`.
  -   When using `nats-js-kv` it is recommended to set `OCIS_CACHE_STORE_NODES` to the same value as `OCIS_EVENTS_ENDPOINT`. That way the cache uses the same nats instance as the event bus.
  -   When using the `nats-js-kv` store, it is possible to set `OCIS_CACHE_DISABLE_PERSISTENCE` to instruct nats to not persist cache data on disc.

## Settings Management

Infinite Scale services can register *settings bundles* with the settings service.

## Settings Usage

Services can set or query Infinite Scale *setting values* of a user from settings bundles.

## Service Accounts

The settings service needs to know the ID's of service accounts but it doesn't need their secrets. Currently only one service account can be configured which has the admin role. This can be set with the `SETTINGS_SERVICE_ACCOUNT_ID_ADMIN` envvar, but it will also pick up the global `OCIS_SERVICE_ACCOUNT_ID` environment variable. Also see the 'auth-service' service description for additional details.

## Default Language

The default language can be defined via the `OCIS_DEFAULT_LANGUAGE` environment variable. If this variable is not defined, English will be used as default. The value has the ISO 639-1 format ("de", "en", etc.) and is limited by the list supported languages. This setting can be used to set the default language for notification and invitation emails.

Important developer note: the list of supported languages is at the moment not easy defineable, as it is the minimum intersection of languages shown in the WebUI and languages defined in the ocis code for the use of notifications and userlog. Even more, not all languages where there are translations available on transifex, are available in the WebUI respectively for ocis notifications, and the translation rate for existing languages is partially not that high. You will see therefore quite often English default strings though a supported language may exist and was selected.

The `OCIS_DEFAULT_LANGUAGE` setting impacts the `notification` and `userlog` services and the WebUI. Note that translations must exist for all named components to be presented correctly.

*   If  `OCIS_DEFAULT_LANGUAGE` **is not set**, the expected behavior is:
    *   The `notification` and `userlog` services and the WebUI use English by default until a user sets another language in the WebUI via _Account -> Language_.
    *    If a user sets another language in the WebUI in _Account -> Language_, then the `notification` and `userlog` services and WebUI use the language defined by the user. If no translation is found, it falls back to English.

*   If  `OCIS_DEFAULT_LANGUAGE` **is set**, the expected behavior is:
    *   The `notification` and `userlog` services and the WebUI use `OCIS_DEFAULT_LANGUAGE`  by default until a user sets another language in the WebUI via _Account -> Language_.
    *   If a user sets another language in the WebUI in _Account -> Language_, the `notification` and `userlog` services and WebUI use the language defined by the user. If no translation is found, it falls back to `OCIS_DEFAULT_LANGUAGE` and then to English.
