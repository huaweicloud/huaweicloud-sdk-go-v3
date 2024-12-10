package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantDto 授权返回数据
type GrantDto struct {

	// 授权id，授权给个人时存在
	Id string `json:"id"`

	// 授权name
	Name *string `json:"name,omitempty"`

	// 授权类型（SECRET，GROUP）
	Type *string `json:"type,omitempty"`

	// 被授权用户信息
	GranteeUser []GrantUserInfo `json:"grantee_user"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o GrantDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantDto struct{}"
	}

	return strings.Join([]string{"GrantDto", string(data)}, " ")
}
