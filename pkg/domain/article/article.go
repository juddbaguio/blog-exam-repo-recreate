package article

type articleService struct {
	DB interface{}
}

type Repository interface {
	CreateArticle(article string) error
}

func NewArticleService(db interface{}) Repository {
	return &articleService{
		DB: db,
	}
}
