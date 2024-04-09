package services

import (
	"EchoSphere/backend/config"
	"EchoSphere/backend/models"
	"EchoSphere/backend/repositories"
)

// TODO: 实现动漫服务的逻辑
type CartoonService struct {
	cartoonRepository *repositories.CartoonRepository
}

// GetCartoonByID 从数据库中获取指定 ID 的动漫信息
func (app *CartoonService) GetCartoonByID() (models.Cartoon, error) {
	// 这里可以是你的数据库操作代码，例如使用 GORM 查询数据库
	var cartoon models.Cartoon
	//if err := config.DB.First(&cartoon, id).Error; err != nil {
	// 读取表 cartoon 中的数据

	if err := config.DB.First(&cartoon).Error; err != nil {
		return cartoon, err
	}
	return cartoon, nil
}

// 读取表 cartoon 中的所有的数据

// 读取表 cartoon 中的所有的数据
func (app *CartoonService) GetAllCartoon() ([]models.Cartoon, error) {
	var cartoons []models.Cartoon
	if err := config.DB.Find(&cartoons).Error; err != nil {
		return cartoons, err
	}
	return cartoons, nil
}
