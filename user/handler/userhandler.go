package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"example.com/go-microservice/user/requests"
	"example.com/go-microservice/user/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//NewHTTPServer - NewHTTPServer
func NewHTTPServer(ctx context.Context, endpoint service.UserEndpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user/save").Handler(httptransport.NewServer(
		endpoint.Save,
		decodeUserSaveReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/find/{id}").Handler(httptransport.NewServer(
		endpoint.Find,
		decodeUserFindReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func decodeUserSaveReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req requests.UserSaveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUserFindReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req requests.UserFindRequest
	vars := mux.Vars(r)

	req = requests.UserFindRequest{
		ID: vars["id"],
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
