package api

import (
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/domain/article"
)

func (s *Server) SetupArticleRoutes() {
	a := article.NewArticleService(s.logic)
	s.Router.HandleFunc("/article", a.CreateArticle).Methods("POST")
	s.Router.HandleFunc("/article/{articleId}", a.ReadArticleById).Methods("GET")

}
