package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProject(router *mux.Router) {

	// list
	router.Handle("/v1/dashboard/projects/metadata", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
		{
			"items": [
				{
					"metadata": {
						"name": "Default",
						"uid": "1"
					}
				},
				{
					"metadata": {
						"name": "TestProject",
						"uid": "2"
					}
				}
			]
		}`))
	}))
}
