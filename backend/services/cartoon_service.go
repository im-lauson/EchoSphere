package services

import (
	"EchoSphere/backend/models"
	"EchoSphere/backend/repositories"
)

// TODO: 实现动漫服务的逻辑
type CartoonService struct {
	cartoonRepository *repositories.CartoonRepository
}

// GetCartoonByID 从数据库中获取指定 ID 的动漫信息
func (cartService *CartoonService) GetCartoonByID(id int) (models.Cartoon, error) {
	// 这里可以是你的数据库操作代码，例如使用 GORM 查询数据库
	var cartoon models.Cartoon
	if err := DB.First(&cartoon, id).Error; err != nil {
		return cartoon, err
	}
	return cartoon, nil
}
