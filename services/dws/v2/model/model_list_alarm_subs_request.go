package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmSubsRequest Request Object
type ListAlarmSubsRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 限制条目数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmSubsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmSubsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmSubsRequest", string(data)}, " ")
}
