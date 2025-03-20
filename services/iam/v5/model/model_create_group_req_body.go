package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupReqBody The request body of create group request.
type CreateGroupReqBody struct {

	// 用户组名，可包含中文、英文、数字、空格、\"_\"、\"-\"、\"{\"和\"}\"的字符串。
	GroupName string `json:"group_name"`

	// 用户组描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`
}

func (o CreateGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupReqBody struct{}"
	}

	return strings.Join([]string{"CreateGroupReqBody", string(data)}, " ")
}
