package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matthewyu246/back/controller"
	"github.com/matthewyu246/back/db"
	"github.com/matthewyu246/back/repository"
	"github.com/matthewyu246/back/usecase"
)

func main() {

	db := db.NewDB()

	photoRepo := repository.NewPhotoRepository(db)
	photoUsecase := usecase.NewPhotoUsecase(photoRepo)
	photoController := controller.NewPhotoController(photoUsecase)

	r := gin.Default()
	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	r.POST("/upload", photoController.UploadPhoto)
	r.GET("/photo/:id", photoController.GetPhoto)

	r.Run(":8080")
}
