package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachAgencyPolicyV5Request Request Object
type AttachAgencyPolicyV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	Body *AttachAgencyPolicyReqBody `json:"body,omitempty"`
}

func (o AttachAgencyPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachAgencyPolicyV5Request struct{}"
	}

	return strings.Join([]string{"AttachAgencyPolicyV5Request", string(data)}, " ")
}
