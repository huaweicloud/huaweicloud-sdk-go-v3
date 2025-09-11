package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionStatisticsBean struct {

	// 活跃SESSION数量
	ActiveSessionNum *int64 `json:"active_session_num,omitempty"`

	// 失败SESSION数量
	FailSessionNum *int64 `json:"fail_session_num,omitempty"`

	// 新增SESSION数量
	NewSessionNum *int64 `json:"new_session_num,omitempty"`

	// 周期
	Period *string `json:"period,omitempty"`
}

func (o SessionStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionStatisticsBean struct{}"
	}

	return strings.Join([]string{"SessionStatisticsBean", string(data)}, " ")
}
