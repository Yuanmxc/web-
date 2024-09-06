package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/studio"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStudio(c *gin.Context) {
	req := &studio.AddStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req.Name)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.AddStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllStudio(c *gin.Context) {
	req := &studio.GetAllStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetAllStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetStudio(c *gin.Context) {
	req := &studio.GetStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UpdateStudio(c *gin.Context) {
	req := &studio.UpdateStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.UpdateStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func DeleteStudio(c *gin.Context) {
	req := &studio.DeleteStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.DeleteStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
