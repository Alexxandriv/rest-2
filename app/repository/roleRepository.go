package repository

import (
	"rest-2/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindRoleById(id int) (dao.Role, error)
	Save(role *dao.Role) (dao.Role, error)
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (r RoleRepositoryImpl) FindRoleById(id int) (dao.Role, error) {
	user := dao.Role{
		ID: id,
	}
	err := r.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find role by id. Error: ", err)
		return dao.Role{}, err
	}
	return user, nil
}

func (r RoleRepositoryImpl) Save(user *dao.Role) (dao.Role, error) {
	var err = r.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save role. Error: ", err)
		return dao.Role{}, err
	}
	return *user, nil
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
