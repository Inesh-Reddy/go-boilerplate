package database

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// What is database pooling ? technique : a certain no.of connections are opened to DB using TCP protocol
// database is also considered as one of the backend
//
//

type Database struct {
	pool *pgxpool.Pool
	log  *zerolog.Logger
}

const DatabasePingTimeout = 10
