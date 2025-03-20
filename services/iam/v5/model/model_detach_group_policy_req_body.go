package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachGroupPolicyReqBody contains information about a urn of a entity
type DetachGroupPolicyReqBody struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`
}

func (o DetachGroupPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachGroupPolicyReqBody struct{}"
	}

	return strings.Join([]string{"DetachGroupPolicyReqBody", string(data)}, " ")
}
