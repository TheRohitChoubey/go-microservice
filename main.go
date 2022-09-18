package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

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

func createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.MkdirAll("./static/vmWare-assignment/src/assets/image/"+albumName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	getAllAlbumsHandler(w, r)
}

func deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.RemoveAll("./static/vmWare-assignment/src/assets/image/" + albumName); err != nil {
		log.Fatal(err)
	}
	getAllAlbumsHandler(w, r)
}

func getAllAlbums(root string) []string {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}

func getAllAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albumNames := getAllAlbums("./static/vmWare-assignment/src/assets/image/")
	response := albumNames
	json.NewEncoder(w).Encode(response)
}

func getAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageNames := getAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func deleteImageFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageName := params["imageName"]

	if err := os.RemoveAll("./static/vmWare-assignment/src/assets/image/" + albumName + "/" + imageName); err != nil {
		log.Fatal(err)
	}
	imageNames := getAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./static/vmWare-assignment/src/assets/image/"+albumName+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	imageNames := getAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/getAllAlbums", getAllAlbumsHandler)
	r.HandleFunc("/createAlbum/{albumName}", createAlbumHandler)
	r.HandleFunc("/deleteAlbum/{albumName}", deleteAlbumHandler)
	r.HandleFunc("/getAllImages/{albumName}", getAllImagesHandler)
	r.HandleFunc("/deleteImageFromAlbum/{albumName}/{imageName}", deleteImageFromAlbumHandler)
	r.HandleFunc("/uploadImage/{albumName}", uploadImageHandler)
	r.HandleFunc("/health", healthHandler)
	log.Println("Server started at :80")
	log.Fatal(http.ListenAndServe(":80", r))
}
