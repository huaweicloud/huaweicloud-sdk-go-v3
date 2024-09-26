package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApproverBasicDto struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 姓名
	Name *string `json:"name,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`

	// 操作类型
	NameCn *string `json:"name_cn,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 头像链接
	AvatarUrl *string `json:"avatar_url,omitempty"`
}

func (o ApproverBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverBasicDto struct{}"
	}

	return strings.Join([]string{"ApproverBasicDto", string(data)}, " ")
}
