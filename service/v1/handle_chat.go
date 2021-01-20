package v1

import "net/http"

func (op Operations) HandleChat(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		op.insertChat(w, r)
	case http.MethodGet:
		op.getChats(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (op Operations) insertChat(w http.ResponseWriter, r *http.Request) {
	// TODO FINISH THIS TO ADD NEW CHATS
}

func (op Operations) getChats(w http.ResponseWriter, r *http.Request) {
	// TODO FINISH THIS TO ADD GET ALL CHATS AND GET CHATS DEPENDING ON THE ROOM
}