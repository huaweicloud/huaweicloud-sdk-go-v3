package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckPolicyInfoResponseInfo 配置检测策略信息
type SecurityCheckPolicyInfoResponseInfo struct {

	// **参数解释**： 策略名称 **取值范围**： 字符长度1-256位
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 策略ID **取值范围**： 字符长度0-64位
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略详情
	Content *string `json:"content,omitempty"`

	PwdPolicyContent *PwdCheckTagInfo `json:"pwd_policy_content,omitempty"`
}

func (o SecurityCheckPolicyInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckPolicyInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckPolicyInfoResponseInfo", string(data)}, " ")
}
