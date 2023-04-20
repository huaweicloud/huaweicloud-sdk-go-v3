package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteDeleteVideoByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteDeleteVideoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeleteVideoByIdResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDeleteVideoByIdResponse", string(data)}, " ")
}
