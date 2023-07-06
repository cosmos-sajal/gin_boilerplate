# gin_boilerplate
A boilerplate for Gin framework, this include the following -
| Integrations Included |
|-----------------------|
| [Connection with DB (PostgreSQL)]([https://example.com/row1](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/database.go#L22)) |
| [Connection with Redis store]([https://example.com/row2](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/redis.go#L17)) |
| [Caching Mechanism with Redis Adapter]([https://example.com/row3](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/helpers/cache_adapter.go)) |

4. A CRUD application. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/controllers/user_controller.go
5. Sentry (error logger) integration.  https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/initialise_error_logger.go
6. Custom Error Logging (to sentry) - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/errorlogger/error_logger.go
7. Request-Response logger middleware. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/logger/logger_middleware.go#L56
8. Dockerised application - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/Dockerfile
9. Setup docker-compose with db, redis, app, workers, brokers and crons - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/docker-compose.yml
10. Cron server setup. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/crons/initialise_cron.go
11. Worker setup with Async Programming. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/connect_async_queue.go
12. JWT auth middleware. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/services/auth/auth_service.go
13. Migrations and Model setup. - https://github.com/cosmos-sajal/gin_boilerplate/blob/main/models/user.go
