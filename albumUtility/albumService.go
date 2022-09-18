package albumUtility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CreateAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.MkdirAll("./static/vmWare-assignment/src/assets/image/"+albumName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	GetAllAlbumsHandler(w, r)
}

func DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.RemoveAll("./static/vmWare-assignment/src/assets/image/" + albumName); err != nil {
		log.Fatal(err)
	}
	GetAllAlbumsHandler(w, r)
}

func GetAllAlbums(root string) []string {
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

func GetAllAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albumNames := GetAllAlbums("./static/vmWare-assignment/src/assets/image/")
	response := albumNames
	json.NewEncoder(w).Encode(response)
}

func GetAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageNames := GetAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func DeleteImageFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageName := params["imageName"]

	if err := os.RemoveAll("./static/vmWare-assignment/src/assets/image/" + albumName + "/" + imageName); err != nil {
		log.Fatal(err)
	}
	imageNames := GetAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
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

	imageNames := GetAllAlbums("./static/vmWare-assignment/src/assets/image/" + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}
