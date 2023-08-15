package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInfoVo 创建信息
type CreateInfoVo struct {

	// 创建时间
	Time *string `json:"time,omitempty"`

	// 创建时间时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`
}

func (o CreateInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInfoVo struct{}"
	}

	return strings.Join([]string{"CreateInfoVo", string(data)}, " ")
}
