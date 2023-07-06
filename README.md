# gin_boilerplate

| Boilerplate Includes |
|-----------------------|
| [Connection with DB (PostgreSQL)](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/database.go#L22) |
| [Connection with Redis store](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/redis.go#L17) |
| [Caching Mechanism with Redis Adapter](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/helpers/cache_adapter.go) |
| [A CRUD application](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/controllers/user_controller.go) |
| [Sentry (error logger) integration](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/initialise_error_logger.go) |
| [Custom Error Logging (to sentry)](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/errorlogger/error_logger.go) |
| [Request-Response logger middleware](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/logger/logger_middleware.go#L56) |
| [Dockerised application](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/Dockerfile) |
| [Setup docker-compose with db, redis, app, workers, brokers and crons](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/docker-compose.yml) |
| [Cron server setup](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/crons/initialise_cron.go) |
| [Worker setup with Async Programming](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/initializers/connect_async_queue.go) |
| [JWT auth middleware](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/services/auth/auth_service.go) |
| [Migrations and Model setup](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/models/user.go) |

### How to setup?
- Take a git clone of the repo.
- Copy the content of `.env.sample` file and create a new file in the root directory as `.env`.
- Change the `JWT_TOKEN` as a new base64 encoded string, you can use below mentioned python code to generate it -
```
import os
import base64

# Generate a random secret key of 32 bytes
random_bytes = os.urandom(32)
secret_key = base64.urlsafe_b64encode(random_bytes).decode('utf-8')

# Print the generated secret key
print(secret_key)
```
- Run `docker-compose build` to build the docker image.
- Run `docker-compose up` to run the docker images. Once they start running, you might need to create a new DB for the repo, for that, do the following -
```
docker ps -> will return all the running containers in your system. Find one by the name "go_boilerplate-db-1"

docker exec -it go_boilerplate-db-1 psql -U postgres -> This will help you exec into the psql server.

create database boilerplate -> This will create the DB named `boilerplate`
```
- Terminate and run `docker-compose up` again, this will run app, worker, cron, db, and redis containers.
