package controllor

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	NestTime  int64   `json:"next_time, omitempty"`
	VideoList []Video `json:"video_list, omitempty"`
}

func Feed(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 
			 FeedResponse{
			 	Response:  Response{
								StatusCode: 0,
								StatusMsg: "hello",
							},
				NestTime:  time.Now().Unix(),
				VideoList: Demovideos,
			})

}
