package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportContainerListRequestBodyCreateTime 创建时间
type ExportContainerListRequestBodyCreateTime struct {

	// 筛选起始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 筛选终止时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ExportContainerListRequestBodyCreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportContainerListRequestBodyCreateTime struct{}"
	}

	return strings.Join([]string{"ExportContainerListRequestBodyCreateTime", string(data)}, " ")
}
