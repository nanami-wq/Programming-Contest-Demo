package router

import (
	"Programming-Demo/core/middleware/web"
	"Programming-Demo/internal/app/File/file_handler"
	"Programming-Demo/internal/app/ai/ai_handler"
	"Programming-Demo/internal/app/user/user_handler"

	"github.com/gin-gonic/gin"
)

func GenerateRouters(r *gin.Engine) *gin.Engine {
	r.GET("/ping", ai_handler.PingMoonshot)
	// 用户相关路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", user_handler.Register)
		userGroup.POST("/login", user_handler.Login)
		userGroup.POST("/logout", web.JWTAuthMiddleware(), user_handler.Logout)
	}
	aiGroup := r.Group("/api/ai")
	{
		aiGroup.POST("/analyze", web.JWTAuthMiddleware(), ai_handler.AnalyzeFile)
		aiGroup.POST("/generation", web.JWTAuthMiddleware(), ai_handler.GenerateLegalDocument)
	}
	// 管理员相关路由
	adminGroup := r.Group("/api/admin", web.JWTAuthMiddleware(), web.AdminAuthMiddleware())
	{
		// 这里可以添加需要管理员权限的路由
		// 管理员可以获取所有用户信息
		adminGroup.GET("/users", user_handler.GetAllUsers)
	}
	fileGroup := r.Group("/api/file", web.JWTAuthMiddleware(), web.AdminAuthMiddleware())
	{
		fileGroup.POST("/upload", file_handler.UploadFileHandler)
		fileGroup.GET("/download/:id", file_handler.DownloadFileHandler)
		fileGroup.DELETE("/delete/:id", file_handler.DeleteFileHandler)
	}
	return r
}
