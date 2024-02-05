package api

import "github.com/google/uuid"
import "github.com/gorilla/mux"

type items struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type Server struct {
	*mux.Router
	shoppingItems []items
}
