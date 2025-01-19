package services

import (
	"gosqllearn/datasource/mysql"
	"gosqllearn/model"
	log "gosqllearn/utils/logging"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDataSamples(c *gin.Context) {
	resp := new(model.Response)

	var sample []model.Sample

	db := mysql.Connect

	result := db.Raw("SELECT id, name FROM sample").Scan(&sample)
	if result.Error != nil {
		go log.Error(result.Error.Error(), "Get data sample error", time.Now())
	}

	if result.RowsAffected < 1 {
		log.Error("no affected rows", "Get data sample mysql no affected rows", time.Now())
		return
	}


	resp.Data = sample
	resp.SetResponse(http.StatusOK, "Get data samples successfuly", "")
	c.JSON(resp.Status, resp)
}
