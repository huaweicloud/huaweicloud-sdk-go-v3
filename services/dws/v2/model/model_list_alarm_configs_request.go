package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmConfigsRequest Request Object
type ListAlarmConfigsRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 限制条目数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmConfigsRequest", string(data)}, " ")
}
