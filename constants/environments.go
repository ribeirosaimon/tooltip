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
	MongoProperties string = "mongo.uri"
	ApiPort         string = "api.port"
)
