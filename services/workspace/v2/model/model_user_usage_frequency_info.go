package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserUsageFrequencyInfo 用户使用频次信息。
type UserUsageFrequencyInfo struct {

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户使用次数。
	UseCount *int32 `json:"use_count,omitempty"`
}

func (o UserUsageFrequencyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserUsageFrequencyInfo struct{}"
	}

	return strings.Join([]string{"UserUsageFrequencyInfo", string(data)}, " ")
}
