package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAntiVirusPolicyResponse Response Object
type CreateAntiVirusPolicyResponse struct {

	// **参数解释**： 任务ID **取值范围**: 字符长度1-64位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**: 策略ID **取值范围**: 字符长度1-64位
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释** 是否全部成功 **取值范围** true: 是 false: 否
	Result *bool `json:"result,omitempty"`

	// **参数解释** 主机结果列表 **取值范围** 不涉及
	FailReasons    *[]FailReasons `json:"fail_reasons,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateAntiVirusPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiVirusPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAntiVirusPolicyResponse", string(data)}, " ")
}
