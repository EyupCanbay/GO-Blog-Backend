package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "golang-blog-backend/models"
	storage "golang-blog-backend/storage"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPost models.Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid JSON post", http.StatusBadRequest)
		return
	}

	newPost.ID = storage.NextID
	storage.NextID++
	storage.Posts = append(storage.Posts, newPost)

	fmt.Println("[POST] New post created", newPost)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "applcication/json")

	err := json.NewEncoder(w).Encode(&storage.Posts)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
