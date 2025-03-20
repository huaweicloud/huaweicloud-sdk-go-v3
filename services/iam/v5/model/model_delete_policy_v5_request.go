package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyV5Request Request Object
type DeletePolicyV5Request struct {

	// 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyV5Request struct{}"
	}

	return strings.Join([]string{"DeletePolicyV5Request", string(data)}, " ")
}
