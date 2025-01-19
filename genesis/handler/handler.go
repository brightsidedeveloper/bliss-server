package handler

import (
	"bliss-server/genesis/queries"
	"bliss-server/genesis/util"
)

type Handler struct {
	Queries *queries.Queries
	JSON    *util.JSON
}
