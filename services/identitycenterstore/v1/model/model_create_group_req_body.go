package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupReqBody 创建用户组请求体
type CreateGroupReqBody struct {

	// 用户组描述
	Description *string `json:"description,omitempty"`

	// 用户组显示名称
	DisplayName string `json:"display_name"`
}

func (o CreateGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupReqBody struct{}"
	}

	return strings.Join([]string{"CreateGroupReqBody", string(data)}, " ")
}
