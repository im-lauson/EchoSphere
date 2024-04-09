package services

import (
	"EchoSphere/backend/config"
	"EchoSphere/backend/models"
)

// TODO: 实现动漫服务的逻辑

// 读取表 cartoon_tencent 中的所有的数据
func GetAllCartoon() ([]models.Cartoon, error) {
	var cartoons []models.Cartoon
	if err := config.DB.Find(&cartoons).Error; err != nil {
		return cartoons, err
	}
	return cartoons, nil
}
