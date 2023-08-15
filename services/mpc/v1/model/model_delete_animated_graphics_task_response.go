package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnimatedGraphicsTaskResponse Response Object
type DeleteAnimatedGraphicsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAnimatedGraphicsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskResponse", string(data)}, " ")
}
