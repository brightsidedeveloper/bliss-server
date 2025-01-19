package handler

import (
	"bliss-server/genesis/queries"
	"bliss-server/genesis/util"
	"bliss-server/injections"
)

type Handler struct {
	Queries    *queries.Queries
	Injections *injections.Injections
	JSON       *util.JSON
}

func NewHandler(queries *queries.Queries, i *injections.Injections, json *util.JSON) *Handler {
	return &Handler{
		Queries:    queries,
		Injections: i,
		JSON:       json,
	}
}
