package internet

import (
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintln(w, "Server is running!")
	})

	mux.HandleFunc("/gobasics", basic)

	if err := http.ListenAndServe(":1910", mux); err != nil {
		fmt.Errorf("Could not run server", err)
	}
}

func basic(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintln(w, "Gobasics page")
}
