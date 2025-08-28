package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserReqBody The request body of update user request.
type UpdateUserReqBody struct {

	// IAM用户名，长度为1到64个字符，只包含字母、数字、\"_\"、\"-\"、\".\"和空格的字符串，且首位不能为数字。
	NewUserName *string `json:"new_user_name,omitempty"`

	// IAM用户描述信息，长度为0到255个字符，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	NewDescription *string `json:"new_description,omitempty"`

	// IAM用户是否启用。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdateUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserReqBody struct{}"
	}

	return strings.Join([]string{"UpdateUserReqBody", string(data)}, " ")
}
