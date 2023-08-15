package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsCallbacksResponse Response Object
type ListVmsCallbacksResponse struct {

	// 请求状态，固定200。
	Status *string `json:"status,omitempty"`

	// 状态描述。
	Message *string `json:"message,omitempty"`

	Data           *ListVmsCallbacksMode `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListVmsCallbacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsCallbacksResponse struct{}"
	}

	return strings.Join([]string{"ListVmsCallbacksResponse", string(data)}, " ")
}
