package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"go-rest-api/confighelper"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Client {
	config := confighelper.GetConfiguration()
	// Set client options
	clientOptions := options.Client().ApplyURI(config.Connection_String)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//collection := client.Database("go_rest_api").Collection("books")

	return client
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {
	fmt.Println("error message========================>", err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, err1 := json.Marshal(response)
	if err1 != nil {
		fmt.Println("error 1=======================================?", err1)
	}

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
