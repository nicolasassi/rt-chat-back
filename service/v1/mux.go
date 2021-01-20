package v1

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"rt-chat/service/v1/operations"
)

const version = "v1"

type Tools struct {
	Client *mongo.Client
}

type Operations struct {
	Rooms *operations.Rooms
}

func NewMux(t *Tools) *http.ServeMux {
	op := new(Operations)
	op.Rooms = operations.NewRooms(t.Client, os.Getenv("MONGO_DBNAME_V1"), os.Getenv("MONGO_ROOMS_COL_V1"))
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("/%s/rooms/", version), op.HandleRooms)
	//mux.HandleFunc("/files", RetrieveFiles)
	return mux
}
