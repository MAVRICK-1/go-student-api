package student

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"github.com/mavrick-1/go-student-api/pkg/types"
)

// New returns a new http.Handler that responds with "Hello, students"
func New() http.HandlerFunc {
    // http.HandlerFunc is an adapter that allows a function with the signature 
    // func(http.ResponseWriter, *http.Request) to be used as an http.Handler.
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // `w` is the http.ResponseWriter interface used to construct the HTTP response.
        // It allows us to write our response data, set headers, etc.
        
        // Examples of methods provided by the http.ResponseWriter interface include:
        // - Write([]byte): Writes the data to the response.
        // - WriteHeader(int): Sends an HTTP response header with the provided status code.
        // - Header(): Allows you to manipulate the headers before sending the response.

        // `r` is the pointer to an http.Request struct that represents the client’s request.
        // It contains all the information about the request, like headers, URL, body, etc.

        // w.Write writes the response body to the http.ResponseWriter.
        // The response body is provided as a byte slice, which is a common type in Go
        // for handling binary and text data.

        var student types.Student

        err:= json.NewDecoder(r.Body).Decode(&student) // Decode the JSON data from the request body into the student struct NewDecoder returns a new decoder that reads from r.
        //Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.

        slog.Info("Student", slog.Any("Student",student))
        if err != nil {
            slog.Error("Error decoding request body", slog.String("Error",err.Error()))
            w.WriteHeader(http.StatusBadRequest)
        }

        // The json.NewDecoder(r.Body).Decode(&student) method reads the JSON data from the request body

        w.Write([]byte("Hello, students"))
    })
}

// The New function returns an http.HandlerFunc that responds with "Hello, students".