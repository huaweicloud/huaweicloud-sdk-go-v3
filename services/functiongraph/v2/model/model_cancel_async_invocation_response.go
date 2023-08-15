package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAsyncInvocationResponse Response Object
type CancelAsyncInvocationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAsyncInvocationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAsyncInvocationResponse struct{}"
	}

	return strings.Join([]string{"CancelAsyncInvocationResponse", string(data)}, " ")
}
