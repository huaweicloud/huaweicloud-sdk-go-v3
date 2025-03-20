package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachUserPolicyReqBody contains information about a urn of a entity
type DetachUserPolicyReqBody struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`
}

func (o DetachUserPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachUserPolicyReqBody struct{}"
	}

	return strings.Join([]string{"DetachUserPolicyReqBody", string(data)}, " ")
}
