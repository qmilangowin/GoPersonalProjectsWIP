package firestoredb

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/qmilangowin/imagebox/pkg/authentication"
	"github.com/qmilangowin/imagebox/pkg/config"
	"github.com/qmilangowin/imagebox/pkg/database/models"
	logging "github.com/qmilangowin/imagebox/pkg/logging"
	"google.golang.org/api/iterator"
)

//default collectionName is posts
var collectionName string = "posts"

var logger *logging.Logger

//PostRepository exposes associated functions and allows access to them
//programatically
type PostRepository interface {
	Save(post *models.Annotation) (*models.Annotation, error)
	GetAll() ([]*models.Annotation, error) //[]models.Annotation
	GetFavouriteQuery() ([]*models.Annotation, error)
	GetLatestPostsQuery() ([]*models.Annotation, error)
}

type repo struct{}

func init() {
	logger = logging.New(config.LogWriter, config.LogDebug)
}
func createClient(ctx context.Context) *firestore.Client {

	gcpAuth := authentication.GetEnvs()
	projectID := gcpAuth.ProjectID

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		logger.PrintErrorf("Failed to create client: %v", err)
	}

	return client
}

//NewPostRepository creates a instance of repo that can be exported
func NewPostRepository() PostRepository {
	return &repo{}
}

//Save saves a new post to Firestore for the given collection
func (*repo) Save(post *models.Annotation) (*models.Annotation, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	if post.CollectionName != "" {
		collectionName = post.CollectionName
	}

	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":          post.ID,
		"Filename":    post.Filename,
		"Description": post.Description,
		"Favourite":   post.Favourite,
		"Created":     time.Now(),
	})

	// _, err := client.Collection(collectionName).Doc("Photos").Set(ctx, map[string]interface{}{
	// 	"ID":          post.ID,
	// 	"Filename":    post.Filename,
	// 	"Description": post.Description,
	// 	"Favourite":   post.Favourite,
	// 	"Created":     time.Now(),
	// }, firestore.MergeAll)

	// _, _, err := client.Collection(collectionName).Doc("Photos").Collection("sub").Add(ctx, map[string]interface{}{
	// 	"ID":          post.ID,
	// 	"Filename":    post.Filename,
	// 	"Description": post.Description,
	// 	"Favourite":   post.Favourite,
	// 	"Created":     time.Now(),
	// })

	// _, err := client.Collection(collectionName).Doc("LA").Delete(ctx)

	if err != nil {
		logger.PrintErrorf("Failed adding a new post to the collection: %v", err)
	}
	return post, nil

}

func (*repo) GetFavouriteQuery() ([]*models.Annotation, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	//TODO: set collection name check via a URL query in GET

	var annotations []*models.Annotation
	query := client.Collection(collectionName).Where("Favourite", "==", true)
	it := query.Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.PrintErrorf("Failed to read from database collection %v", err)
			return nil, err
		}

		annotation := models.Annotation{
			ID: doc.Data()["ID"].(int64),
			//CollectionName: doc.Data()["CollectionName"].(string),
			Filename:    doc.Data()["Filename"].(string),
			Description: doc.Data()["Description"].(string),
			//Filepath:    doc.Data()["Filepath"].(string),
			Favourite: doc.Data()["Favourite"].(bool),
			Created:   doc.Data()["Created"].(time.Time),
		}
		annotations = append(annotations, &annotation)
	}

	return annotations, nil

}

func (*repo) GetLatestPostsQuery() ([]*models.Annotation, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var annotations []*models.Annotation

	query := client.Collection(collectionName).OrderBy("Created", firestore.Desc).Limit((5))
	it := query.Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.PrintErrorf("Failed to read from database collection %v", err)
			return nil, err
		}

		annotation := models.Annotation{
			ID: doc.Data()["ID"].(int64),
			//CollectionName: doc.Data()["CollectionName"].(string),
			Filename:    doc.Data()["Filename"].(string),
			Description: doc.Data()["Description"].(string),
			//Filepath:    doc.Data()["Filepath"].(string),
			Favourite: doc.Data()["Favourite"].(bool),
			Created:   doc.Data()["Created"].(time.Time),
		}
		annotations = append(annotations, &annotation)
	}

	return annotations, nil

}

//GetAll returns all posts/annotations from the Firestore DB
func (*repo) GetAll() ([]*models.Annotation, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	//TODO: set collection name check via a URL query in GET

	var annotations []*models.Annotation

	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.PrintErrorf("Failed to read from database collection %v", err)
			return nil, err
		}

		annotation := models.Annotation{
			ID: doc.Data()["ID"].(int64),
			//CollectionName: doc.Data()["CollectionName"].(string),
			Filename:    doc.Data()["Filename"].(string),
			Description: doc.Data()["Description"].(string),
			//Filepath:    doc.Data()["Filepath"].(string),
			Favourite: doc.Data()["Favourite"].(bool),
			Created:   doc.Data()["Created"].(time.Time),
		}
		annotations = append(annotations, &annotation)
	}

	return annotations, nil
}
