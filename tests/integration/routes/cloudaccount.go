package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCloudAccount(router *mux.Router) {

	// create
	router.Handle("/v1/cloudaccounts/aws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"uid": "1"}`))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))

	// get and delete
	router.Handle("/v1/cloudaccounts/aws/{uid}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`
			{
				"apiVersion": "v1",
				"kind": "AwsCloudAccount",
				"metadata": {
					"name": "aws-cloud-account",
					"annotations": {
						"scope": "project"
					},
					"uid": "1"
				},
				"spec": {
					"credentialType": "secret",
					"accessKey": "fake-access-key",
					"partition": "aws"
				},
				"status": {
					"state": "active"
				}
			}`))
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))

	// validate
	router.Handle("/v1/clouds/aws/account/validate", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}))
}
