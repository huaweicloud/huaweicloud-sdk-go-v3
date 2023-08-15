package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionAsyncInvokeConfigResponse Response Object
type DeleteFunctionAsyncInvokeConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFunctionAsyncInvokeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
