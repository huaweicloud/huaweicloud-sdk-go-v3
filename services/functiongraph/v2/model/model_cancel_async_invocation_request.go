package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAsyncInvocationRequest Request Object
type CancelAsyncInvocationRequest struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	Body *CancelAsyncInvocationRequestBody `json:"body,omitempty"`
}

func (o CancelAsyncInvocationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAsyncInvocationRequest struct{}"
	}

	return strings.Join([]string{"CancelAsyncInvocationRequest", string(data)}, " ")
}
