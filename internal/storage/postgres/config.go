// Package postgres provides a postgres storage implementation for managing data in a PostgreSQL database.
package postgres

// Config contains postgres configuration.
type Config struct {
	URI                   string `long:"uri" env:"URI" description:"PGX connection uri to the postgres" default:"postgresql://pg:test@localhost:5432/pgw?sslmode=disable" required:"true" yaml:"uri"`
	Schema                string `long:"schema" env:"SCHEMA" description:"target pg schema" default:"public"`
	DisableSimpleProtocol bool   `long:"simple.protocol"  description:"disable implicit prepared statement usage (if PreferSimpleProtocol == false wit bouncer usage it will produce errors in prepared statements)" `
}
