package apis

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"models"
    db "databases"
	"net/http"
	"strconv"
)

func AddBannerApi(c *gin.Context) {

	var p models.ScBanner
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

func GetBannerAllApi(c *gin.Context) {

	ra, err := models.ScBanners(db.SqlDB)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"banner": ra,
	})
}

func GetBannerApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		id = 0
	}

	p, err := models.ScBannerByID(db.SqlDB, uint(id))

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"banner": p,
	})
}

