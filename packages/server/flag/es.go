package flag

import (
	"web/models"
)

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
