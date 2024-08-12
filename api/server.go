package api

import db "github.com/nmhoang2909/bank/db/sqlc"

type Server struct {
	store *db.Store
}
