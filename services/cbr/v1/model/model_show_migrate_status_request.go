package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMigrateStatusRequest struct {

	// 是否查询其他区域结果
	AllRegions *bool `json:"all_regions,omitempty"`
}

func (o ShowMigrateStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrateStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrateStatusRequest", string(data)}, " ")
}
