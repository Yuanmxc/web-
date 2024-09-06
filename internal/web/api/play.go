package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/play"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPlay(c *gin.Context) {
	req := &play.AddPlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.AddPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func UpdatePlay(c *gin.Context) {
	req := &play.UpdatePlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.UpdatePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func DeletePlay(c *gin.Context) {
	req := &play.DeletePlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.DeletePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllPlay(c *gin.Context) {
	req := &play.GetAllPlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetAllPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func AddSchedule(c *gin.Context) {
	req := &play.AddScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.AddSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func UpdateSchedule(c *gin.Context) {
	req := &play.UpdateScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.UpdateSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func DeleteSchedule(c *gin.Context) {
	req := &play.DeleteScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.DeleteSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllSchedule(c *gin.Context) {
	req := &play.GetAllScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetAllSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
