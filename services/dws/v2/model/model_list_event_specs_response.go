package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventSpecsResponse struct {

	// 事件配置总数
	Count *int32 `json:"count,omitempty"`

	// 事件配置列表
	EventSpecs     *[]EventSpecResponse `json:"event_specs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListEventSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSpecsResponse struct{}"
	}

	return strings.Join([]string{"ListEventSpecsResponse", string(data)}, " ")
}
