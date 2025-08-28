package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Group 用户组。
type Group struct {

	// 用户组ID。
	GroupId string `json:"group_id"`

	// 用户组名，长度为1到128个字符，可包含中文、英文、数字、空格、\"_\"、\"-\"、\"{\"和\"}\"的字符串。
	GroupName string `json:"group_name"`

	// 用户组创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 统一资源名称。
	Urn string `json:"urn"`

	// 用户组描述信息，长度为0到255字符，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	Description *string `json:"description,omitempty"`
}

func (o Group) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Group struct{}"
	}

	return strings.Join([]string{"Group", string(data)}, " ")
}
