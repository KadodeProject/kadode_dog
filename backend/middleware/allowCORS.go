package middleware

import(
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
)

func AllowCORS(r *gin.Engine) *gin.Engine {
	/**
    * CORS関連
    */
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "https://kadode.usuyuki.net",
            "https://kado.day",
            "https://wiki.kado.day",
            "https://diary.kado.day",
            "https://api.kado.day",
            "https://portal.kado.day",
            "https://paper.kado.day",
            "http://localhost:7000",
            "http://localhost:7001",
            "http://localhost:7002",
            "http://localhost:7003",
            "http://localhost:7004",
            "http://localhost:7005",
        },
        AllowMethods: []string{
            "POST",
            "GET",
        },
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
        },
        AllowCredentials: true,
        MaxAge: 24 * time.Hour,
    }))
	return r
}
