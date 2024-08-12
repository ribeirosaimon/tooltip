package constants

type env string

const (
	AERGIA env = "AERGIA_ENV"
)

const (
	DEV  env = "dev"
	PROD env = "prod"
	TEST env = "test"
)

const (
	PostgressUrl      string = "postgress.url"
	PostgressDatabase string = "postgress.database"
	MongoUrl          string = "mongo.url"
	MongoDatabase     string = "mongo.database"
	ApiPort           string = "api.port"
)
