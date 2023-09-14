package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionUrlResponse Response Object
type DeleteFunctionUrlResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionUrlResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionUrlResponse", string(data)}, " ")
}
