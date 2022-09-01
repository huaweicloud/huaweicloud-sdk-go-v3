package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMoreInstantMessagesResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 留言列表
	MessageList    *[]QueryMessageInfoV2 `json:"message_list,omitempty" xml:"message_list"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMoreInstantMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMoreInstantMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListMoreInstantMessagesResponse", string(data)}, " ")
}
