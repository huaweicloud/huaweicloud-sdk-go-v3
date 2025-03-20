package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupReqBody The request body of update group request.
type UpdateGroupReqBody struct {

	// 用户组名，可包含中文、英文、数字、空格、\"_\"、\"-\"、\"{\"和\"}\"的字符串。
	NewGroupName *string `json:"new_group_name,omitempty"`

	// 用户组描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	NewGroupDescription *string `json:"new_group_description,omitempty"`
}

func (o UpdateGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupReqBody struct{}"
	}

	return strings.Join([]string{"UpdateGroupReqBody", string(data)}, " ")
}
