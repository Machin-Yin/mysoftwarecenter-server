package main

import (
	. "apis"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/banner", AddBannerApi)
	router.GET("/banner/:id", GetBannerApi)
	router.GET("/banners", GetBannerAllApi)

	router.POST("/recommend", AddRecommendApi)
	router.GET("/recommend/:id", GetRecommendApi)
	router.GET("/recommends", GetRecommendAllApi)

	router.POST("/user", AddUserApi)
	router.PUT("/user/:id", ModUserApi)
	router.GET("/user/:id", GetUserApi)
	router.GET("/users", GetUsersAllApi)

	router.POST("/comment", AddCommentApi)
	router.POST("/comments", GetCommentsApi)
	router.GET("/comment/:id", GetCommentApi)

	router.POST("/release", AddReleaseApi)
	router.POST("/releases", GetReleasesApi)
	router.GET("/release/:id", GetReleaseApi)

	router.POST("/product", AddProductApi)
	router.PUT("/product/:id", ModProductApi)
	router.GET("/product/:id", GetProductApi)
	router.GET("/products", GetProductAllApi)
	router.GET("/product_from_name", GetProductFromNameApi)

	router.POST("/category", AddCategoryApi)
	router.PUT("/category/:id", ModCategoryApi)
	router.GET("/category/:id", GetCategoryApi)
	router.GET("/categories", GetCategoryAllApi)

	router.StaticFS("/icon", http.Dir("./test/icon_url"))

	return router
}
