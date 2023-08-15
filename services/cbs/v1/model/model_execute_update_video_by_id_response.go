package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUpdateVideoByIdResponse Response Object
type ExecuteUpdateVideoByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteUpdateVideoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUpdateVideoByIdResponse struct{}"
	}

	return strings.Join([]string{"ExecuteUpdateVideoByIdResponse", string(data)}, " ")
}
