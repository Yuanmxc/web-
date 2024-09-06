package main

import (
	"TTMS/internal/web/api"
	"TTMS/internal/web/mw"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(mw.LimitMiddleware())

	baseGroup := r.Group("/ttms")

	baseGroup.POST("/user/create", api.CreateUser)
	baseGroup.POST("/user/login", api.UserLogin)
	baseGroup.POST("/user/verify", api.GetVerification)
	baseGroup.POST("/user/forget", api.ForgetPassword)
	//以上内容不需要Token鉴权
	baseGroup.GET("/user/all", mw.AuthMiddleware(), api.GetAllUser)
	baseGroup.POST("/user/change", mw.AuthMiddleware(), api.ChangeUserPassword)
	baseGroup.POST("/user/delete", mw.AuthMiddleware(), api.DeleteUser)
	baseGroup.GET("/user/info", mw.AuthMiddleware(), api.GetUserInfo)
	baseGroup.POST("/user/bind", mw.AuthMiddleware(), api.BindEmail)

	studioGroup := baseGroup.Group("/studio", mw.AuthMiddleware())
	studioGroup.POST("/add", api.AddStudio)
	studioGroup.GET("/all", api.GetAllStudio)
	studioGroup.GET("/info", api.GetStudio)
	studioGroup.POST("/update", api.UpdateStudio)
	studioGroup.POST("/delete", api.DeleteStudio)

	seatGroup := baseGroup.Group("/seat", mw.AuthMiddleware())
	seatGroup.POST("/add", api.AddSeat)
	seatGroup.POST("/update", api.UpdateSeat)
	seatGroup.POST("/delete", api.DeleteSeat)
	seatGroup.GET("/all", api.GetAllSeat)

	playGroup := baseGroup.Group("/play", mw.AuthMiddleware())
	playGroup.POST("/add", api.AddPlay)
	playGroup.POST("/update", api.UpdatePlay)
	playGroup.POST("/delete", api.DeletePlay)
	playGroup.GET("/all", api.GetAllPlay)

	scheduleGroup := baseGroup.Group("/schedule", mw.AuthMiddleware())
	scheduleGroup.POST("/add", api.AddSchedule)
	scheduleGroup.POST("/update", api.UpdateSchedule)
	scheduleGroup.POST("/delete", api.DeleteSchedule)
	scheduleGroup.GET("/all", api.GetAllSchedule)

	ticketGroup := baseGroup.Group("/ticket", mw.AuthMiddleware())
	ticketGroup.POST("/update", api.UpdateTicket)
	ticketGroup.GET("/all", api.GetAllTicket)
	ticketGroup.POST("/buy", api.BuyTicket)
	ticketGroup.POST("/return", api.ReturnTicket)

	orderGroup := baseGroup.Group("/order", mw.AuthMiddleware())
	orderGroup.POST("/commit", api.CommitOrder)
	orderGroup.GET("/all", api.GetAllOrder)
	orderGroup.GET("/analysis", api.GetOrderAnalysis)
	return r
}
