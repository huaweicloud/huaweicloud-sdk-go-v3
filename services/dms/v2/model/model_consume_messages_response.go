package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ConsumeMessagesResponse struct {

	// 消息数组。
	Body           *[]ConsumeMessage `json:"body,omitempty" xml:"body"`
	HttpStatusCode int               `json:"-"`
}

func (o ConsumeMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeMessagesResponse struct{}"
	}

	return strings.Join([]string{"ConsumeMessagesResponse", string(data)}, " ")
}
