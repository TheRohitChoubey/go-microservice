package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	albumUtility "github.com/TheRohitChoubey/go-microservice/albumUtility"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking Application Health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving HomePage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is Up and Running")
}

func openHtmlView(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	r := mux.NewRouter()
	tpl, _ = tpl.ParseGlob("client/*.html")
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/client", openHtmlView)
	r.HandleFunc("/getAllAlbums", albumUtility.GetAllAlbumsHandler)
	r.HandleFunc("/createAlbum/{albumName}", albumUtility.CreateAlbumHandler)
	r.HandleFunc("/deleteAlbum/{albumName}", albumUtility.DeleteAlbumHandler)
	r.HandleFunc("/getAllImages/{albumName}", albumUtility.GetAllImagesHandler)
	r.HandleFunc("/deleteImageFromAlbum/{albumName}/{imageName}", albumUtility.DeleteImageFromAlbumHandler)
	r.HandleFunc("/getImageFromAlbum/{albumName}/{imageName}", albumUtility.GetImageFromAlbumHandler)
	r.HandleFunc("/uploadImage/{albumName}", albumUtility.UploadImageHandler)
	r.HandleFunc("/health", healthHandler)
	log.Println("Server started at :80")
	log.Fatal(http.ListenAndServe(":80", r))
}
