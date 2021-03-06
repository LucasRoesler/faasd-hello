package function

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	log.Printf("received request: %s\n", string(req.Body))

	message := fmt.Sprintf("Hello world, input was: %s", string(req.Body))
	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
