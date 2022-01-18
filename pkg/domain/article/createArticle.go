package article

import (
	"encoding/json"
	"net/http"

	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/models"
)

func (a *articleService) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var payload models.Article
	jsonResponseEncoder := json.NewEncoder(w)
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		jsonResponseEncoder.Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	if errResp := payload.Validate(); errResp != nil {
		jsonResponseEncoder.Encode(errResp)

		return
	}

	articleId, err := a.logic.Create(models.Article{})

	if err != nil {
		jsonResponseEncoder.Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	jsonResponseEncoder.Encode(articleId)

}
