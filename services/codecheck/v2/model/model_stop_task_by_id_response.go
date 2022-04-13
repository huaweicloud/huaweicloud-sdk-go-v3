package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopTaskByIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopTaskByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTaskByIdResponse struct{}"
	}

	return strings.Join([]string{"StopTaskByIdResponse", string(data)}, " ")
}
