package article

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *articleService) ReadArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonResponseEncoder := json.NewEncoder(w)

	if vars != nil {
		jsonResponseEncoder.Encode(map[string]string{
			"error": "no articleId found",
		})

		return
	}

	articleId, err := strconv.Atoi(vars["articleId"])

	if err != nil {
		jsonResponseEncoder.Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	articleSlice, err := a.logic.Read(articleId)

	if err != nil {
		jsonResponseEncoder.Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	jsonResponseEncoder.Encode(articleSlice)
}
