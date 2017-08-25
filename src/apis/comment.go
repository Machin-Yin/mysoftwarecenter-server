package apis

import (
	db "databases"
	"log"
	"models"
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func AddCommentApi(c *gin.Context) {

	var p models.ScComment
	if err := c.Bind(&p); err != nil {
		log.Println(err)
		return
	}

	err := p.Save(db.SqlDB)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": p.ID,
	})
}

func GetCommentApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Println(err)
		id = 0
	}

	p, err := models.ScCommentByID(db.SqlDB, uint(id))

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": p,
	})
}

func GetCommentsApi(c *gin.Context) {
    var ids []uint
	err := c.Bind(&ids)
	if err != nil {
		log.Println(err)
		return
	}
    
    srs, err := models.ScCommentByIDs(db.SqlDB, ids)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"releases": srs,
	})
}

