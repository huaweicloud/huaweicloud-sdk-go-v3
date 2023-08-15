package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncInvokeFunctionResponse Response Object
type AsyncInvokeFunctionResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AsyncInvokeFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncInvokeFunctionResponse struct{}"
	}

	return strings.Join([]string{"AsyncInvokeFunctionResponse", string(data)}, " ")
}
