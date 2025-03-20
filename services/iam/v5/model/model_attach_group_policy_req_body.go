package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachGroupPolicyReqBody Contains information about a id of a group.
type AttachGroupPolicyReqBody struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`
}

func (o AttachGroupPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachGroupPolicyReqBody struct{}"
	}

	return strings.Join([]string{"AttachGroupPolicyReqBody", string(data)}, " ")
}
