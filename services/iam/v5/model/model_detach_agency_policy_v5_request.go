package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachAgencyPolicyV5Request Request Object
type DetachAgencyPolicyV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	Body *DetachAgencyPolicyReqBody `json:"body,omitempty"`
}

func (o DetachAgencyPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachAgencyPolicyV5Request struct{}"
	}

	return strings.Join([]string{"DetachAgencyPolicyV5Request", string(data)}, " ")
}
