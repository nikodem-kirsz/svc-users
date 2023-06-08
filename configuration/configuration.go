package configuration

type Config struct {
	DB                    DB
	StatusServer          StatusServer
	ReadinessProbleServer StatusServer
	Log                   Logging
	EventProducer         EventProducer
	Metrics               Metrics
	JWT                   JWT
}

type DB struct {
	UseCertificates   bool   `env:"DB_USE_CERTS"`
	CA                string `env:"DB_CA"`
	Cert              string `env:"DB_CERT"`
	Key               string `env:"DB_KEY"`
	User              string `env:"DB_USER,required"`
	Password          string `env:"DB_PASSWORD"`
	Host              string `env:"DB_HOST,required"`
	Name              string `env:"DB_NAME" envDefault:"users"`
	Port              int    `env:"DB_PORT" envDefault:"3306"`
	Instance          string `env:"DB_INSTANCE"`
	ProjectID         string `env:"GCP_PROJECT_ID"`
	EnableGormLogMode bool   `env:"ENABLE_GORM_LOGMODE" envDefault:"false"`
}

type StatusServer struct {
}

type Logging struct {
}

type EventProducer struct {
}

type Metrics struct {
}

type JWT struct {
}

func Load(*Config) {

}
