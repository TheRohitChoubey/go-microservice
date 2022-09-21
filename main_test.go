package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	albumUtility "github.com/TheRohitChoubey/go-microservice/albumUtility"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

func TestFunction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Handler Testing")
}

var _ = Describe("AlbumUtility.albumService", func() {

	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When get request is sent to root path", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				rootHandler,
			)
		})

		It("Returns the default message", func() {
			log.Println("Test1 is running")
			resp, err := http.Get(server.URL() + "/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))

			body, bodyErr := ioutil.ReadAll(resp.Body)
			response := string(body)
			Expect(bodyErr).ShouldNot(HaveOccurred())
			Expect(response).Should(Equal("Application is Up and Running"))
		})
	})

	Context("When get request is sent to getAllAlbums", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				albumUtility.GetAllAlbumsHandler,
				albumUtility.CreateAlbumHandler,
			)
		})

		It("Returns the string array of album names", func() {
			log.Println("Test2 is running")

			createAlbumRequest, err := http.NewRequest(http.MethodGet, server.URL()+"/createAlbum/{albumName}", nil)
			createAlbumResponse := httptest.NewRecorder()
			vars := map[string]string{
				"albumName": "testAlbum",
			}
			createAlbumRequest = mux.SetURLVars(createAlbumRequest, vars)

			albumUtility.CreateAlbumHandler(createAlbumResponse, createAlbumRequest)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(createAlbumResponse.Result().StatusCode).Should(Equal(http.StatusOK))

			getAlbumRequest, getAlbumErr := http.NewRequest(http.MethodGet, server.URL()+"/getAllAlbums", nil)
			getAlbumResponse := httptest.NewRecorder()

			albumUtility.GetAllAlbumsHandler(getAlbumResponse, getAlbumRequest)

			getAllAlbumResponse := strings.Split(getAlbumResponse.Body.String(), "\"")

			albumNameArray := make([]string, 0, len(getAllAlbumResponse)/2+1)

			for index, element := range getAllAlbumResponse {
				if index%2 != 0 {
					albumNameArray = append(albumNameArray, element)
				}
			}

			Expect(getAlbumErr).ShouldNot(HaveOccurred())
			Expect(getAlbumResponse.Result().StatusCode).Should(Equal(http.StatusOK))
			Expect(len(albumNameArray)).ShouldNot(Equal(0))
		})
	})

	Context("When get request is sent to createAlbum", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				albumUtility.CreateAlbumHandler,
			)
		})

		It("Returns the string array of album names with new album name in it", func() {
			log.Println("Test3 is running")
			createAlbumRequest, err := http.NewRequest(http.MethodGet, server.URL()+"/createAlbum/{albumName}", nil)
			createAlbumResponse := httptest.NewRecorder()
			vars := map[string]string{
				"albumName": "testAlbum",
			}
			createAlbumRequest = mux.SetURLVars(createAlbumRequest, vars)

			albumUtility.CreateAlbumHandler(createAlbumResponse, createAlbumRequest)

			getAllAlbumResponse := strings.Split(createAlbumResponse.Body.String(), "\"")

			albumNameArray := make([]string, 0, len(getAllAlbumResponse)/2+1)

			for index, element := range getAllAlbumResponse {
				if index%2 != 0 {
					albumNameArray = append(albumNameArray, element)
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(createAlbumResponse.Result().StatusCode).Should(Equal(http.StatusOK))
			Expect(len(albumNameArray)).ShouldNot(Equal(0))
		})
	})

	Context("When get request is sent to deleteAlbum", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				albumUtility.DeleteAlbumHandler,
				albumUtility.CreateAlbumHandler,
			)
		})

		It("Returns the string array of album names with new album deleted from it", func() {
			log.Println("Test4 is running")
			createAlbumRequest, err := http.NewRequest(http.MethodGet, server.URL()+"/createAlbum/{albumName}", nil)
			createAlbumResponse := httptest.NewRecorder()
			vars := map[string]string{
				"albumName": "testAlbum",
			}
			createAlbumRequest = mux.SetURLVars(createAlbumRequest, vars)

			albumUtility.CreateAlbumHandler(createAlbumResponse, createAlbumRequest)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(createAlbumResponse.Result().StatusCode).Should(Equal(http.StatusOK))

			getAllAlbumResponse := strings.Split(createAlbumResponse.Body.String(), "\"")

			albumNameArray := make([]string, 0, len(getAllAlbumResponse)/2+1)

			for index, element := range getAllAlbumResponse {
				if index%2 != 0 {
					albumNameArray = append(albumNameArray, element)
				}
			}

			sizeAfterCreation := len(albumNameArray)

			deleteAlbumRequest, deleteAlbumErr := http.NewRequest(http.MethodGet, server.URL()+"/deleteAlbum/{albumName}", nil)
			deleteAlbumResponse := httptest.NewRecorder()

			deleteAlbumRequest = mux.SetURLVars(deleteAlbumRequest, vars)

			albumUtility.DeleteAlbumHandler(deleteAlbumResponse, deleteAlbumRequest)

			Expect(deleteAlbumErr).ShouldNot(HaveOccurred())
			Expect(deleteAlbumResponse.Result().StatusCode).Should(Equal(http.StatusOK))

			afeterDeleteAllAlbumResponse := strings.Split(deleteAlbumResponse.Body.String(), "\"")

			afterDeleteAlbumNameArray := make([]string, 0, len(afeterDeleteAllAlbumResponse)/2+1)

			for index, element := range afeterDeleteAllAlbumResponse {
				if index%2 != 0 {
					afterDeleteAlbumNameArray = append(afterDeleteAlbumNameArray, element)
				}
			}

			sizeAfterDeletion := len(afterDeleteAlbumNameArray)
			Expect(sizeAfterCreation).Should(Equal(sizeAfterDeletion + 1))

		})
	})

})
