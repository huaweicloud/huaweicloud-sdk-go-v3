package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserReqBody The request body of create user request.
type CreateUserReqBody struct {

	// IAM用户名，长度为1到64个字符，只包含字母、数字、\"_\"、\"-\"、\".\"和空格的字符串，且首位不能为数字。
	Name string `json:"name"`

	// IAM用户描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`

	// IAM用户是否启用。
	Enabled bool `json:"enabled"`
}

func (o CreateUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserReqBody struct{}"
	}

	return strings.Join([]string{"CreateUserReqBody", string(data)}, " ")
}
