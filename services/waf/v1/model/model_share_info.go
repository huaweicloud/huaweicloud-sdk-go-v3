package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShareInfo 共享信息
type ShareInfo struct {

	// 共享用户数
	ShareCount *int32 `json:"share_count,omitempty"`

	// 接受用户数
	AcceptCount *int32 `json:"accept_count,omitempty"`

	// 处理状态
	ProcessStatus *int32 `json:"process_status,omitempty"`
}

func (o ShareInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareInfo struct{}"
	}

	return strings.Join([]string{"ShareInfo", string(data)}, " ")
}
