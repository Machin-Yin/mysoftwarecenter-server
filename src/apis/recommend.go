package apis

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"models"
    db "databases"
	"net/http"
	"strconv"
)

func AddRecommendApi(c *gin.Context) {

	var p models.ScRecommend
	err := c.Bind(&p)
	if err != nil {
		log.Println(err)
		return
	}

	err = p.Save(db.SqlDB)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": p.ID,
	})
}

func GetRecommendAllApi(c *gin.Context) {

	ra, err := models.ScRecommends(db.SqlDB)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recommend": ra,
	})
}

func GetRecommendApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		id = 0
	}

	p, err := models.ScRecommendByID(db.SqlDB, uint(id))

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recommend": p,
	})
}

