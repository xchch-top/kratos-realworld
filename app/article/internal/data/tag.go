package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/app/article/internal/biz"
	pkgData "kratos-realworld/pkg/data"
)

type Tag struct {
	pkgData.Model
	Name     string
	Articles []Article
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
