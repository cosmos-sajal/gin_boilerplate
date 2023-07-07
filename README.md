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

### How to use the repository?
#### Adding a controller
- Create a controller function in the `controllers` package, for example -
```
func SignInController(c *gin.Context) {
	var requestBody validators.SignInStruct
	var randomOTP string
	c.Bind(&requestBody)
	validationErr := requestBody.Validate()
	if validationErr != nil {
		c.JSON(validationErr.Status, validationErr)
		return
	}

	user, _ := models.GetUserByMobile(*requestBody.MobileNumber)
	cacheKey := otpservice.OTP_KEY_PREFIX + user.MobileNumber
	val, _ := helpers.GetCacheValue(cacheKey)
	if val == "" {
		randomOTP = helpers.GenerateRandomOTP()
		helpers.SetCacheValue(cacheKey, randomOTP, 120)
	} else {
		randomOTP = val
	}

	otpservice.SendOTP(user.MobileNumber, randomOTP, int(user.ID))

	c.JSON(http.StatusOK, gin.H{
		"result": "OTP Sent successfully",
	})
}
```
- Create validation in the validators package, here you will validate the request body and params. e.g. - [validator](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/validators/auth_validator.go)

#### Models & Migrations
- Check [Gorm Documenation](https://gorm.io/docs/) & [sql-migrate](https://github.com/rubenv/sql-migrate).
- Add the models like [this](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/models/user.go)
- You can either run [Automigrate](https://gorm.io/docs/migration.html) functionality provided by gorm to migrate the DB or use sql-migrate to update.
##### How to automigrate
Run the following from the root folder after creating the model
```
go run migrate/migrate.go
```
##### How to run migration by creating SQL migrations using sql-migrate
- Read the documentation provided by [sql-migrate](https://github.com/rubenv/sql-migrate)
- Run the following command to create an empty migration with up/down
```
docker-compose run --rm app sh -c 'sql-migrate new -config=dbconfig.yml -env="development" <migration_name>'

<migration_name> can be anything, it is just for readability, this will be used to name the SQL file.
```
- The SQL file is created under the directory `migrations` (this is specified in dbconfig.yml under dir key)
- The file will be empty, you need to write the SQL commands for migration, example -
```
-- +migrate Up
CREATE INDEX idx_mobile_number ON public.users USING btree (mobile_number);

-- +migrate Down
DROP INDEX idx_mobile_number;
```
- Run `docker-compose run --rm app sh -c 'sql-migrate up -config=dbconfig.yml -env="development" create_index_mobile_no'` to run the up migration.
- Run `docker-compose run --rm app sh -c 'sql-migrate down -config=dbconfig.yml -env="development"'` for down migration.
- You can also check the status of the current migration using this command - `docker-compose run --rm app sh -c 'sql-migrate status -config=dbconfig.yml -env="development"'`, this will return something like this -
```
+-------------------------------------------+---------+
|                 MIGRATION                 | APPLIED |
+-------------------------------------------+---------+
| 20230705135358-init.sql                   | no      |
| 20230705141125-create_index_mobile_no.sql | no      |
+-------------------------------------------+---------+
```
This output comes from gorp-migrations table from the DB.

#### How to use Redis Cache
- Check for a file named [cache_adapter.go](https://github.com/cosmos-sajal/gin_boilerplate/blob/main/helpers/cache_adapter.go)
- Use this to interact with Redis.
- Example:
```
GetCacheValue
val, _ := helpers.GetCacheValue(cacheKey)
if val == "" {
	randomOTP = helpers.GenerateRandomOTP()
	helpers.SetCacheValue(cacheKey, randomOTP, 120)
} else {
	randomOTP = val
}

SetCacheValue
helpers.SetCacheValue(key, "1", OTP_ATTEMPT_KEY_EXPIRY)
```


