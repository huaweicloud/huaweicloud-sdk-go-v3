package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogRetentionCommonSettingsResponse Response Object
type ShowLogRetentionCommonSettingsResponse struct {
	LogRetention *CommonSettingsResponseLogRetention `json:"log_retention,omitempty"`

	// 日志存储磁盘占用上限
	DataUsageLimit *int32 `json:"data_usage_limit,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowLogRetentionCommonSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogRetentionCommonSettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowLogRetentionCommonSettingsResponse", string(data)}, " ")
}
