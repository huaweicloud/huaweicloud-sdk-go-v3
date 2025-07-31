package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupPolicyRequestBody 待删除的安全组策略
type DeleteSecurityGroupPolicyRequestBody struct {

	// 待删除的安全组策略
	Policies []DeletedPolicy `json:"policies"`
}

func (o DeleteSecurityGroupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupPolicyRequestBody", string(data)}, " ")
}
