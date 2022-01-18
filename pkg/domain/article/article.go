package article

import (
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/logic"
)

type articleService struct {
	logic logic.Logic
}

func NewArticleService(logic logic.Logic) *articleService {
	return &articleService{
		logic: logic,
	}
}
