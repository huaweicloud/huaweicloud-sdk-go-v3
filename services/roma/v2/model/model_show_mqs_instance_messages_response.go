package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMqsInstanceMessagesResponse struct {

	// 消息列表。
	Messages *[]ShowMqsInstanceMessagesRespMessages `json:"messages,omitempty" xml:"messages"`

	// 消息总数。
	Total float32 `json:"total,omitempty" xml:"total"`

	// 返回总条数。
	Size           float32 `json:"size,omitempty" xml:"size"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMqsInstanceMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceMessagesResponse", string(data)}, " ")
}
