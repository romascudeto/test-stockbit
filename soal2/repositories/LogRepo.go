package repositories

import (
	"soal2/config"
	"soal2/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateLog(log *models.Log) (err error) {

	if err = config.DB.Create(log).Error; err != nil {
		return err
	}
	return nil
}
