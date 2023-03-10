package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/app/user/internal/biz"
	pkgData "kratos-realworld/pkg/data"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	pkgData.Model
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

func (u *User) TableName() string {
	return "user"
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	user := User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
		Model:        *pkgData.NewCreateModel(),
	}
	rv := r.data.db.Create(&user)
	if rv.Error != nil {
		return nil, rv.Error
	}
	return &biz.User{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := new(User)
	result := r.data.db.Where("email = ?", email).First(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:           u.Id,
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserById(ctx context.Context, id uint64) (*biz.User, error) {
	var user User
	result := r.data.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.User{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		Bio:          user.Bio,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, bizUser *biz.User) error {
	user := User{
		Username: bizUser.Username,
		Email:    bizUser.Email,
		Bio:      bizUser.Bio,
		Image:    bizUser.Image,
		Model:    *pkgData.NewUpdateModel(bizUser.Id),
	}
	result := r.data.db.Select("username", "email", "bio", "image", "updated_time").Updates(&user)

	return result.Error
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	var user User
	result := r.data.db.Where("username = ?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.User{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
