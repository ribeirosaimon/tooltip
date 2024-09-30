package constants

type EnvironmentType string

const (
	AERGIA EnvironmentType = "AERGIA_ENV"
)

const (
	DEV         EnvironmentType = "dev"
	PROD        EnvironmentType = "prod"
	INTEGRATION EnvironmentType = "integration_test"
	TEST        EnvironmentType = "test"
)

const (
	PostgressUrl      string = "postgress.url"
	PostgressDatabase string = "postgress.database"
	MongoUrl          string = "mongo.url"
	MongoDatabase     string = "mongo.database"
	ApiPort           string = "api.port"

	// SecretKey this block is about security
	SecretKey      string = "secretKey"
	LoggedUser     string = "loggedUser"
	PasswordJwtKey string = "password.jwt.key"
)
