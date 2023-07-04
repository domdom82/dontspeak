package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080"
	}

	instanceIndex, found := os.LookupEnv("CF_INSTANCE_INDEX")
	if !found {
		panic("expected CF_INSTANCE_INDEX env var")
	}

	if len(os.Args) < 2 {
		panic("expected 1 argument for CF_INSTANCE_INDEX")
	}

	targetIndex := os.Args[1]

	fmt.Println("CF_INSTANCE_INDEX", instanceIndex, "TARGET_INDEX", targetIndex)

	if targetIndex == instanceIndex {
		startListening(port)
	} else {
		waitSilently()
	}

}

func waitSilently() {
	for {
		fmt.Println("waiting forever silently...")
		time.Sleep(1 * time.Minute)
	}
}

func startListening(port string) {
	fmt.Println("start listening...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
