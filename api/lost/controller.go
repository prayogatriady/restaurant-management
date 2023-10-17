package lost

import (
	"fmt"
	"net/http"

	"github.com/prayogatriady/restaurant-management/http/httpresponse"

	"github.com/gin-gonic/gin"
)

func LostInSpace(c *gin.Context) {

	var errorMessage string
	errorMessage = "You are lost in space"

	if lang := c.GetHeader("lang"); lang == "ID" {
		errorMessage = "Anda tersesat"
	}

	httpresponse.BaseResponse(&httpresponse.HttpParams{
		GinContext:  c,
		StatusCode:  http.StatusOK,
		ServiceName: "LostInSpace",
		Errors:      fmt.Sprintf("No Route: %s", errorMessage),
	})
}
