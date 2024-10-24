package main // Declares the package name, "main", as this is the entry point for the application

import (
	"context" // Import "context" for managing the lifecycle of the application
	"fmt"     // Import "fmt" for formatted I/O operations
	"log"     // Import "log" for logging errors
	"log/slog"
	"net/http"  // Import "net/http" for HTTP server and request handling
	"os"        // Import "os" for working with the operating system
	"os/signal" // Import "os/signal" for handling signals
	"syscall"   // Import "syscall" for system calls
	"time"      // Import "time" for working with time

	"github.com/mavrick-1/go-student-api/pkg/config" // Import the configuration package from the local module
)

func main() {
	// Load the configuration using the MustLoad function from the config package
	cfg := config.MustLoad() // This function loads the configuration from a file or environment variables

	// Setup a new HTTP request multiplexer (router)
	router := http.NewServeMux() // Creates a new ServeMux for routing HTTP requests
	// serverMux is an HTTP request multiplexer that matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL

	// Define the root route ("/") and its handler function
	// The handler function responds to HTTP requests with "Hello, students"
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, students") // Write the response "Hello, students" to the ResponseWriter
	})

	slog.Info("Starting server on", slog.String("Address",cfg.Address)) // Log a message indicating that the server is starting and the address it is listening on
	//slog.String is used to create a structured log field with a key-value pair eg
	// slog.String("Address", cfg.Address) creates a structured log field with the key "Address" and the value of cfg.Address
	//slog.Info logs an informational message

	//slog is a structured logger that provides a way to log structured data in a consistent format


	// Set up the HTTP server with the address and the router
	server := http.Server{
		Addr:    cfg.HTTPServer.Address, // Address where the server listens (e.g., localhost:8080)
		Handler: router,                 // The ServeMux router handles incoming HTTP requests
	}

	done := make(chan os.Signal, 1) // Create a channel to receive signals
	// make is a built-in function that creates a new channel with a buffer size of 1\
	// channels are used to communicate between goroutines
	// goroutines are lightweight threads that are managed by the Go runtime
	// buffer means that the channel can hold one signal at a time
	// os.Signal is a type that represents an operating system signal (e.g., interrupt, termination)
	// A channel is a communication mechanism that allows one goroutine to send a message to another goroutine

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // Notify the channel when an interrupt signal is received (e.g., Ctrl+C)
	// params: channel, signals to listen for (e.g., interrupt, SIGINT, SIGTERM) SIGINT is the interrupt signal (Ctrl+C), SIGTERM is the termination signal, and os.Interrupt is a generic interrupt signal

	go func() {
		// Start the server by calling ListenAndServe, which listens on the specified address
		err := server.ListenAndServe() // This function blocks and runs the server until an error occurs

		// If an error occurs when starting the server, log the error and exit the program
		if err != nil {
			log.Fatal(err) // Logs the error and stops the program if the server fails to start
		}
	}()

	<-done // Wait for a signal to be received on the channel

	slog.Info("Shutting down server...") // Log a message indicating that the server is shutting down

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Create a context with a timeout of 5 seconds to shut down the server
	// context.WithTimeout creates a new context with a timeout that will be canceled after the specified duration, 
	// context.Background() returns an empty context that has no values associated with it and is never canceled,
	// 5*time.Second is the duration of the timeout
	// cancel is a function that can be called to cancel the context
	//ctx is a context that can be used to manage the lifecycle of the application

	fmt.Println(ctx) //context.Background.WithDeadline(2024-10-25 00:04:50.2806882 +0530 IST m=+8.024178601 [5s]) // it will print the context with the deadline
	fmt.Println(cancel) // it will print the address of the cancel function //0x4f4b20
	defer cancel() // Defer the cancel function to ensure it is called before the function returns
	//explain defer
	// defer is used to ensure that a function call is performed later in a program's execution, usually for cleanup operations
	// defer is often used to close files, release resources, or unlock mutexes
	//eg defer file.Close() // Close the file when the function returns

	err := server.Shutdown(ctx) // Shutdown the server gracefully using the context
	// server.Shutdown gracefully shuts down the server without interrupting any active connections
	// it takes a context as an argument to allow for a timeout or cancellation of the shutdown process

	if err != nil {
		slog.Error("Server shutdown error:", slog.String("error",err.Error())) // Log an error message if the server shutdown encounters an error
	}

	slog.Info("Server stopped") // Log a message indicating that the server has stopped


}