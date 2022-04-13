package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteFunctionTriggersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteFunctionTriggersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFunctionTriggersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteFunctionTriggersResponse", string(data)}, " ")
}
