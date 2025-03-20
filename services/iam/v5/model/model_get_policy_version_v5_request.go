package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPolicyVersionV5Request Request Object
type GetPolicyVersionV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`

	// 身份策略版本号，以\"v\"开头后跟数字的字符串，例如\"v5\"。
	VersionId string `json:"version_id"`
}

func (o GetPolicyVersionV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPolicyVersionV5Request struct{}"
	}

	return strings.Join([]string{"GetPolicyVersionV5Request", string(data)}, " ")
}
