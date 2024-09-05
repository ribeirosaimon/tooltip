package constants

type EnvironmentType string

const (
	AERGIA EnvironmentType = "AERGIA_ENV"
)

const (
	DEV  EnvironmentType = "dev"
	PROD EnvironmentType = "prod"
	TEST EnvironmentType = "test"
)

const (
	PostgressUrl      string = "postgress.url"
	PostgressDatabase string = "postgress.database"
	MongoUrl          string = "mongo.url"
	MongoDatabase     string = "mongo.database"
	ApiPort           string = "api.port"
)
