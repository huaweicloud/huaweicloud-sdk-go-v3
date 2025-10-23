package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportSettingsRequest Request Object
type ListReportSettingsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 偏移量
	Offset *int64 `json:"offset,omitempty"`

	// 单次请求限制数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListReportSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportSettingsRequest struct{}"
	}

	return strings.Join([]string{"ListReportSettingsRequest", string(data)}, " ")
}
