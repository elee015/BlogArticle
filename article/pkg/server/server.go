package server

import (
	"context"
	"encoding/json"
	"net/http"

	articlepb "github.com/blog/article/pkg/pb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/articles").Handler(httptransport.NewServer(
		endpoints.AddArticle,
		decodeArticleReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/articles/{id}").Handler(httptransport.NewServer(
		endpoints.GetArticle,
		decodeGetArticleReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/articles").Handler(httptransport.NewServer(
		endpoints.ListArticles,
		decodeListArticleReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "applicaiton/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeArticleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req articlepb.AddArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetArticleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req articlepb.GetArticleRequest

	vars := mux.Vars(r)

	req = articlepb.GetArticleRequest{
		Id: vars["id"],
	}

	return req, nil
}

func decodeListArticleReq(ctx context.Context, r *http.Request) (interface{}, error) {
	req := articlepb.ListArticlesRequest{}

	return req, nil
}
