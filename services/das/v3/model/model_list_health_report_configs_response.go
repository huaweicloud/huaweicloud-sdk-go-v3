package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHealthReportConfigsResponse Response Object
type ListHealthReportConfigsResponse struct {

	// 实例配置详情列表。
	Configs        *[]InstanceHealthReportConfigDto `json:"configs,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListHealthReportConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListHealthReportConfigsResponse", string(data)}, " ")
}
