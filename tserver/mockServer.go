package tserver

type MockEnvironment struct {
	PgsqlHost       string
	PgsqlDatabase   string
	PgsqlEntryPoint string

	RedisHost       string
	RedisDatabase   string
	RedisEntryPoint string
}

func NewMockEnvironment(mock MockEnvironment) {
	env.Config.HostName = "https://test.io"
	env.Pgsql = dbConfig{
		Host:       mock.PgsqlHost,
		Database:   mock.PgsqlDatabase,
		EntryPoint: mock.PgsqlEntryPoint,
	}
	env.Redis = dbConfig{
		Host:       mock.RedisHost,
		Database:   mock.RedisDatabase,
		EntryPoint: mock.RedisEntryPoint,
	}
}
