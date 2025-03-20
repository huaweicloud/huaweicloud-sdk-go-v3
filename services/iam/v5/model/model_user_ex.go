package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserEx IAM用户。
type UserEx struct {

	// IAM用户描述信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`

	// IAM用户名，长度为1到64个字符，只包含字母、数字、\"_\"、\"-\"、\".\"和空格的字符串，且首位不能为数字。
	UserName string `json:"user_name"`

	// IAM用户是否为根用户。
	IsRootUser bool `json:"is_root_user"`

	// IAM用户创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	// 统一资源名称。
	Urn string `json:"urn"`

	// IAM用户是否启用。
	Enabled bool `json:"enabled"`

	// 自定义标签列表。
	Tags []Tag `json:"tags"`
}

func (o UserEx) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserEx struct{}"
	}

	return strings.Join([]string{"UserEx", string(data)}, " ")
}
