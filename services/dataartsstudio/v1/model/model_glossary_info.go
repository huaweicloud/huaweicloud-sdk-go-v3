package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlossaryInfo struct {

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 标签的guid
	Guid *string `json:"guid,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 创建时间
	CreateTime float32 `json:"create_time,omitempty"`
}

func (o GlossaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlossaryInfo struct{}"
	}

	return strings.Join([]string{"GlossaryInfo", string(data)}, " ")
}
