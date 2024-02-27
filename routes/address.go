package routes

import (
	"github.com/gin-gonic/gin"
	"llm_gateway/config"
	"llm_gateway/model"
	"net/http"
)

type UpdateAddrReq struct {
	//Addr string `json:"addr" binding:"required"`
	Type string `json:"type" binding:"required"`
}

func readUserIp(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func GetAddress(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: config.GetServerConfig()})
}
func UpdateAddress(c *gin.Context) {
	var req UpdateAddrReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ip := readUserIp(c.Request)
	if req.Type == config.T2I {
		config.UpdateT2IAddr(ip)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "啊对对对" + ip})
}
