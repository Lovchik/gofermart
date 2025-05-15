package storage

type Storage interface {
	HealthCheck() (err error)
	IsUserExists() (err error)
}
