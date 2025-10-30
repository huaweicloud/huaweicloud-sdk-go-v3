package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultSecurityCheckPolicyResponse Response Object
type ShowDefaultSecurityCheckPolicyResponse struct {

	// **参数解释**： 策略名称 **取值范围**： 字符长度1-256位
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 策略ID **取值范围**： 字符长度0-64位
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略详情
	Content *string `json:"content,omitempty"`

	PwdPolicyContent *PwdCheckTagInfo `json:"pwd_policy_content,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowDefaultSecurityCheckPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultSecurityCheckPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultSecurityCheckPolicyResponse", string(data)}, " ")
}
