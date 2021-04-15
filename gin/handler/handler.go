package handler

import (
	"ceph/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/alertmanager/template"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) (result *Handler) {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PrometheusSend(c *gin.Context) {

	data := template.Data{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}
	fmt.Println(data)
	if err := h.service.Send(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}
