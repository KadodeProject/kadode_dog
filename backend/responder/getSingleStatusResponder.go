package responder

import(
	"github.com/gin-gonic/gin"
	"backend/typefile"
)
func getSingleStatusResponder(c *gin.Context,data typefile.SingleStatusResponseType) {
	c.JSON(200, gin.H{
		"id": data.Id,
		"service_id": data.Service_id,
		"status": data.Status,
		"response_ms": data.Response_ms,
	})
}

