package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"time"

	"github.com/qmilangowin/imagebox/pkg/authentication"

	"math/rand"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/qmilangowin/imagebox/pkg/config"
	firestoredb "github.com/qmilangowin/imagebox/pkg/database/firestore"
	"github.com/qmilangowin/imagebox/pkg/database/models"
	logging "github.com/qmilangowin/imagebox/pkg/logging"
	"github.com/qmilangowin/imagebox/pkg/web"
)

var db firestoredb.PostRepository = firestoredb.NewPostRepository()
var logger *logging.Logger

func init() {
	logger = logging.New(config.LogWriter, config.LogDebug)
}

func (app *ServerApplication) home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Now Allowed", 405)
		return
	}

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	latest, err := db.GetLatestPostsQuery()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.render(w, r, "home.page.html", data)

	// app.render replaces below - leaving below in for legacy reasons
	// files := []string{
	// 	"./ui/html/home.page.html",
	// 	"./ui/html/base.layout.html",
	// 	"./ui/html/footer.partial.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	http.Error(w, "Internal Server Error", 500)
	// }

}

func (app *ServerApplication) getGalleries(w http.ResponseWriter, r *http.Request) {

	all, err := db.GetAll()
	if err != nil {
		app.Log.PrintErrorf("could not fetch from database %s", err)
	}
	app.render(w, r, "galleries.page.html", &web.TemplateData{Images: all})

}

func (app *ServerApplication) getSingleImage(w http.ResponseWriter, req *http.Request) {
	//TODO: decide on query to fetch images

	posts, err := db.GetAll()
	if err != nil {
		app.ServerError(w, err)
	}

	files := []string{
		"./ui/html/showimage.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.Execute(w, posts)
	if err != nil {
		app.ServerError(w, err)
	}

}

func (app *ServerApplication) getFaves(w http.ResponseWriter, r *http.Request) {

	faves, err := db.GetFavouriteQuery()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	data := &web.TemplateData{Images: faves}

	app.render(w, r, "showfaves.page.html", data)

}

func (app *ServerApplication) uploadImage(resp http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		resp.Header().Set("Allow", http.MethodPost)
		app.ClientError(resp, http.StatusMethodNotAllowed)
		return
	}

	var post models.Annotation
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		app.ServerError(resp, err)
		return
	}

	gcpBucket := authentication.GetEnvs()
	bucket := gcpBucket.Bucket
	fileName := "notes.txt"
	uploads := "/Users/milan-macbook-air/Documents/Repos/PersonalProjects/ImageBox/test_files"
	object := path.Join(uploads, fileName)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.PrintErrorf("Error occurred setting storage client %s", err)

	}
	defer client.Close()

	//open local file
	f, err := os.Open(path.Join(uploads, fileName))
	fmt.Println(object)
	if err != nil {
		logger.PrintErrorf("Could not open file for upload: %s", err)
		return
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	//upload object with storage writer
	wc := client.Bucket(bucket).Object(fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, f); err != nil {
		logger.PrintError(err)
	}
	if err := wc.Close(); err != nil {
		logger.PrintError(err)
	}

	//logger.Infoprintf("file %s uploaded ", object)
	app.Log.Infoprintf("file %s uploaded ", fileName)

	//write upload details to Firestore
	post.ID = rand.Int63()
	db.Save(&post)
	app.replyHeaders(resp, req, http.StatusOK)
	json.NewEncoder(resp).Encode(post)

}
