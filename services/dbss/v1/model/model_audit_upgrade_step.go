package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditUpgradeStep struct {

	// 消息
	Msg *string `json:"msg,omitempty"`

	// 结果码
	ResultCode *string `json:"result_code,omitempty"`

	// 步骤名称
	StepName *string `json:"step_name,omitempty"`

	// 升级时间
	Time *string `json:"time,omitempty"`
}

func (o AuditUpgradeStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditUpgradeStep struct{}"
	}

	return strings.Join([]string{"AuditUpgradeStep", string(data)}, " ")
}
