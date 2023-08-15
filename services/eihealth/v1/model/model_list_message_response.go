package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageResponse Response Object
type ListMessageResponse struct {

	// 消息总数
	Count *int32 `json:"count,omitempty"`

	// 消息列表
	Messages       *[]MessageRsp `json:"messages,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageResponse struct{}"
	}

	return strings.Join([]string{"ListMessageResponse", string(data)}, " ")
}
