package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/ticket"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdateTicket(c *gin.Context) {
	req := &ticket.UpdateTicketRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}

	resp, err := rpc.UpdateTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}

func GetAllTicket(c *gin.Context) {
	req := &ticket.GetAllTicketRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetAllTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}

func BuyTicket(c *gin.Context) {
	req := &ticket.BuyTicketRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	UserId, _ := c.Get("ID")
	req.UserId = UserId.(int64)
	resp, err := rpc.BuyTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}

func ReturnTicket(c *gin.Context) {
	req := &ticket.ReturnTicketRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	UserId, _ := c.Get("ID")
	req.UserId = UserId.(int64)
	resp, err := rpc.ReturnTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
}
