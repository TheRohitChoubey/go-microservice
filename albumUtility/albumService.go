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

var rootPath = "/images/"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func CreateAlbumHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("logs in CreateAlbumHandler")
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.MkdirAll(rootPath+albumName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	GetAllAlbumsHandler(w, r)
}

func DeleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("logs in DeleteAlbumHandler")
	params := mux.Vars(r)
	albumName := params["albumName"]
	if err := os.RemoveAll(rootPath + albumName); err != nil {
		log.Fatal(err)
	}
	GetAllAlbumsHandler(w, r)
}

func GetAllAlbums(root string) []string {
	var files []string
	log.Println("logs in GetAllAlbums")
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
	enableCors(&w)

	log.Println("Inside GetAllAlbumsHandler")
	curdir, err := os.Getwd()

	//check if any error occurs
	if err != nil {
		//display error if
		fmt.Println(err)
	}

	//display the path
	log.Println(curdir)

	albumNames := GetAllAlbums(rootPath)
	response := albumNames
	json.NewEncoder(w).Encode(response)
}

func GetAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	albumName := params["albumName"]
	log.Println("logs in GetAllImagesHandler")
	imageNames := GetAllAlbums(rootPath + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func DeleteImageFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("logs in DeleteImageFromAlbumHandler")
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageName := params["imageName"]

	if err := os.RemoveAll(rootPath + albumName + "/" + imageName); err != nil {
		log.Fatal(err)
	}
	imageNames := GetAllAlbums(rootPath + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}

func GetImageFromAlbumHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	albumName := params["albumName"]
	imageName := params["imageName"]

	img, err := os.Open(rootPath + albumName + "/" + imageName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("logs in GetImageFromAlbumHandler")
	defer img.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+imageName)
	w.Header().Set("Content-Type", "application/force-download")
	io.Copy(w, img)
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	f, err := os.OpenFile(rootPath+albumName+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	log.Println("logs in UploadImageHandler", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	imageNames := GetAllAlbums(rootPath + albumName)
	response := imageNames
	json.NewEncoder(w).Encode(response)
}
