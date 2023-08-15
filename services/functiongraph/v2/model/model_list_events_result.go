package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsResult 获取测试事件响应返回体。
type ListEventsResult struct {

	// 测试事件ID。
	Id *string `json:"id,omitempty"`

	// 上次修改的时间。
	LastModified float32 `json:"last_modified,omitempty"`

	// 测试事件名称。
	Name *string `json:"name,omitempty"`
}

func (o ListEventsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResult struct{}"
	}

	return strings.Join([]string{"ListEventsResult", string(data)}, " ")
}
