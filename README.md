# dp-topic-api

Enables greater flexibility in creating journeys through the website

## Getting started

* Run `make debug`

### Dependencies

To run the Topic API locally requires the following:

* Mongo db (you can use [dp-compose](https://github.com/ONSdigital/dp-compose) to stand up an instance in a local Docker container)
* Once you have a working mongo db instance, you will want to populate your database with topics - see `./scripts/README.md` for seeding scripts
* No further dependencies other than those defined in `go.mod`. However, although by default the Topic API has the ENABLE_PRIVATE_ENDPOINTS environment variable set to false, note that it is commonly set to true locally (in the .zshrc - for use in the dp-compose stacks) so it may need to be explicitly set to false:

```shell
export ENABLE_PRIVATE_ENDPOINTS=false
```

Otherwise, AWS Cognito authorisation will be needed, e.g. run [the auth stack in dp-compose](https://github.com/ONSdigital/dp-compose/tree/main/v2/stacks/auth#auth-stack) and then log into your local Florence using your sandbox Florence credentials.

### Tools

To run some of our tests you will need additional tooling:

#### Audit

We use `dis-vulncheck` to do auditing, which you will [need to install](https://github.com/ONSdigital/dis-vulncheck).

#### Linting

We use v2 of golangci-lint, which you will [need to install](https://golangci-lint.run/docs/welcome/install).

#### Validating Specification

To run `make validate-specification`, which can be used to validate the swagger spec, you require Node v20.x and to install @redocly/cli:

```shell
npm install -g @redocly/cli
```

### Configuration

| Environment variable         | Default                                           | Description                                                                                                        |
|------------------------------|---------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| BIND_ADDR                    | :25300                                            | The host and port to bind to                                                                                       |
| GRACEFUL_SHUTDOWN_TIMEOUT    | 10s                                               | The graceful shutdown timeout in seconds (`time.Duration` format)                                                  |
| HEALTHCHECK_INTERVAL         | 30s                                               | Time between self-healthchecks (`time.Duration` format)                                                            |
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s                                               | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format) |
| MONGODB_BIND_ADDR            | localhost:27017                                   | The MongoDB bind address                                                                                           |
| MONGODB_USERNAME             |                                                   | MongoDB Username                                                                                                   |
| MONGODB_PASSWORD             |                                                   | MongoDB Password                                                                                                   |
| MONGODB_DATABASE             | topics                                            | The MongoDB topics database                                                                                        |
| MONGODB_COLLECTIONS          | TopicsCollection:topics,ContentCollection:content | MongoDB collections                                                                                                |
| MONGODB_ENABLE_READ_CONCERN  | false                                             | Switch to use (or not) majority read concern                                                                       |
| MONGODB_ENABLE_WRITE_CONCERN | true                                              | Switch to use (or not) majority write concern                                                                      |
| MONGODB_CONNECT_TIMEOUT      | 5s                                                | The timeout when connecting to MongoDB (`time.Duration` format)                                                    |
| MONGODB_QUERY_TIMEOUT        | 15s                                               | The timeout for querying MongoDB (`time.Duration` format)                                                          |
| MONGODB_IS_SSL               | false                                             | Switch to use (or not) TLS when connecting to mongodb                                                              |
| ZEBEDEE_URL                  | http://localhost:8082                             | The URL to Zebedee (for authentication)                                                                            |
| ENABLE_PRIVATE_ENDPOINTS     | false                                             | Enable private endpoints for the API                                                                               |
| ENABLE_PERMISSIONS_AUTHZ     | false                                             | Enable/disable user/service permissions checking for private endpoints                                             |

## Environments

Any data issues in any of the ONS environments, please visit [dp-operations](https://github.com/ONSdigital/dp-operations) repository for guides

* [Updating Release Date and Publishing Topic Guide](https://github.com/ONSdigital/dp-operations/blob/main/data-fixes/update-topic-release-date.md#update-topic-release-date)
* [Update Topic slug](https://github.com/ONSdigital/dp-operations/blob/main/data-fixes/add-slug-to-topics.md)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics <https://www.ons.gov.uk>

Released under MIT license, see [LICENSE](LICENSE.md) for details.
