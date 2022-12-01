package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警统计
type AlarmStatisticResponse struct {

	// 日期
	Date *string `json:"date,omitempty"`

	// 紧急
	Urgent *string `json:"urgent,omitempty"`

	// 重要
	Important *string `json:"important,omitempty"`

	// 次要
	Minor *string `json:"minor,omitempty"`

	// 提示
	Prompt *string `json:"prompt,omitempty"`
}

func (o AlarmStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmStatisticResponse struct{}"
	}

	return strings.Join([]string{"AlarmStatisticResponse", string(data)}, " ")
}
