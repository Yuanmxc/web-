package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/order"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAllOrder(c *gin.Context) {
	req := &order.GetAllOrderRequest{}

	UserId, _ := c.Get("ID")
	req.UserId = UserId.(int64)
	t := c.Query("OrderType")
	orderType, err := strconv.Atoi(t)
	if err != nil {
		log.Println("err = ", err, "orderType = ", orderType)
	}
	if orderType == 2 {
		req.OrderType = 2
	}
	resp, err := rpc.GetAllOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetOrderAnalysis(c *gin.Context) {
	req := &order.GetOrderAnalysisRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetOrderAnalysis(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func CommitOrder(c *gin.Context) {
	req := &order.CommitOrderRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	UserId, _ := c.Get("ID")
	req.UserId = UserId.(int64)
	resp, err := rpc.CommitOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
