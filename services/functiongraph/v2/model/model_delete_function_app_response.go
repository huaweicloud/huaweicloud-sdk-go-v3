package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionAppResponse Response Object
type DeleteFunctionAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAppResponse", string(data)}, " ")
}
