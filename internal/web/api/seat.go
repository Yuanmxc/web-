package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/studio"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSeat(c *gin.Context) {
	req := &studio.AddSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.AddSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllSeat(c *gin.Context) {
	req := &studio.GetAllSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.GetAllSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UpdateSeat(c *gin.Context) {
	req := &studio.UpdateSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.UpdateSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
func DeleteSeat(c *gin.Context) {
	req := &studio.DeleteSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.DeleteSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
