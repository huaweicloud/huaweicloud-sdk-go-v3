package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimCallbacksResponse Response Object
type ListAimCallbacksResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	// 回调地址列表。
	Data           *[]Callback `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListAimCallbacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimCallbacksResponse struct{}"
	}

	return strings.Join([]string{"ListAimCallbacksResponse", string(data)}, " ")
}
