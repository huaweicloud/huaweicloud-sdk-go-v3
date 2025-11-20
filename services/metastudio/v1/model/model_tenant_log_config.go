package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantLogConfig 租户日志配置
type TenantLogConfig struct {

	// 日志组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流id
	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o TenantLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantLogConfig struct{}"
	}

	return strings.Join([]string{"TenantLogConfig", string(data)}, " ")
}
