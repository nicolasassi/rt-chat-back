package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rt-chat/service/v1/operations"
)

func (op Operations) HandleRooms(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		op.insertRoom(w, r)
	case http.MethodDelete:
		op.deleteRoom(w, r)
	case http.MethodGet:
		op.getRooms(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (op Operations) insertRoom(w http.ResponseWriter, r *http.Request) {
	var m map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error while reading body: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("error while unmarshaling body: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	if name, ok := m["name"].(string); ok {
		roomsEntity, err := op.Rooms.GetRoomByName(r.Context(), name)
		if err != nil {
			log.Printf("error while retrieving room %s: %v", name, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if roomsEntity != nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
		if err := op.Rooms.InsertNewRoom(r.Context(), name); err != nil {
			log.Printf("error while inserting room %s: %v", name, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
			log.Printf("error while responding: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	return
}

func (op Operations) deleteRoom(w http.ResponseWriter, r *http.Request) {
	if name := r.URL.Query().Get("name"); name != "" {
		if err := op.Rooms.DeleteRoomByName(r.Context(), name); err != nil {
			log.Printf("error while deleting room %s: %v", name, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
			log.Printf("error while responding: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, http.StatusText(http.StatusAccepted), http.StatusAccepted)
	return
}

func (op Operations) getRooms(w http.ResponseWriter, r *http.Request) {
	var rooms []operations.RoomEntity
	if name := r.URL.Query().Get("name"); name != "" {
		roomEntity, err := op.Rooms.GetRoomByName(r.Context(), name)
		if err != nil {
			log.Printf("error while retrieving room %s: %v", name, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if roomEntity == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		rooms = append(rooms, *roomEntity)
	} else {
		roomsEntity, err := op.Rooms.GetRooms(r.Context())
		if err != nil {
			log.Printf("error while retrieving rooms %s: %v", name, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if roomsEntity == nil {
			http.Error(w, http.StatusText(http.StatusAccepted), http.StatusAccepted)
			return
		}
		for _, room := range *roomsEntity {
			rooms = append(rooms, room)
		}
	}
	b, err := json.Marshal(rooms)
	if err != nil {
		log.Printf("error while marshaling rooms %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Printf("error while responding: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	return
}
