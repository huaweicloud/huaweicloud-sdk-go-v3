package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnSharedInfo 共享连接信息
type ConnSharedInfo struct {

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 共享连接创建时间
	SharedTime *int64 `json:"shared_time,omitempty"`

	// 共享连接过期时间
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o ConnSharedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnSharedInfo struct{}"
	}

	return strings.Join([]string{"ConnSharedInfo", string(data)}, " ")
}
