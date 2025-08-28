package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// User IAM用户。
type User struct {

	// IAM用户描述信息，长度为0到255个字符，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`

	// IAM用户名，长度为1到64个字符，只包含字母、数字、\"_\"、\"-\"、\".\"和空格的字符串，且首位不能为数字。
	UserName string `json:"user_name"`

	// IAM用户是否为根用户。
	IsRootUser bool `json:"is_root_user"`

	// IAM用户创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// IAM用户ID。
	UserId string `json:"user_id"`

	// 统一资源名称。
	Urn string `json:"urn"`

	// IAM用户是否启用。
	Enabled bool `json:"enabled"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
