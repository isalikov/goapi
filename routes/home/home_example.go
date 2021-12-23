package home

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/isalikov/goapi/internal/services"
	"github.com/isalikov/goapi/internal/utils"
	proto "github.com/isalikov/goservice/pkg/api/grpc"
	"net/http"
	"time"
)

func Example(ctx context.Context, clients *services.Clients, logger *utils.Logger) func (c *gin.Context)  {
	// swagger:operation GET /home/example Home Example
	//
	// Example
	//
	// ---
	// produces:
	// - application/json
	//
	// responses:
	//   '201':
	//     description: successful
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Aircraft"
	//
	//   '400':
	//     description: Invalid Request
	//
	//   '500':
	//     description: Server Error
	return func(c *gin.Context) {
		res, err := clients.GoService.ServiceMethod(ctx, &proto.ServiceMethodRequest{
			Message: "...",
		})
		if err != nil {
			t := time.Now()

			logger.Log(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Server Error",
				"time": t,
			})

			return
		}

		c.JSON(http.StatusOK, res.Message)
	}
}
