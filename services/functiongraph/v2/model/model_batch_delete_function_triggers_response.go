package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFunctionTriggersResponse Response Object
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
