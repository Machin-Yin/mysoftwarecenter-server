package apis

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"models"
    db "databases"
	"net/http"
	"strconv"
)

func AddUserApi(c *gin.Context) {

	var p models.ScUser
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

func GetUsersAllApi(c *gin.Context) {

	ra, err := models.ScUsers(db.SqlDB)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": ra,
	})
}

func GetUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		id = 0
	}

	p, err := models.ScUserByID(db.SqlDB, uint(id))

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": p,
	})
}

func ModUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		id = 0
	}

	var p models.ScUser
	err = c.Bind(&p)
	if err != nil {
		log.Println(err)
		return
	}

	p.ID = uint(id)

	err = p.Save(db.SqlDB)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":   p.ID,
	})
}

