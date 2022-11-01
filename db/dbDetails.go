package db

import (
	"github.com/jmoiron/sqlx"
)

type DbDetails struct {
	PokemonDb       *sqlx.DB
	UsePokemonCache bool
	GeneralDb       *sqlx.DB
}