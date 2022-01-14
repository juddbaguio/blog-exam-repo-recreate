package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) SetupArticleRoutes() {
	s.Router.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"message": s.Article.CreateArticle("test arg").Error(),
		})
	}).Methods("GET")

}
