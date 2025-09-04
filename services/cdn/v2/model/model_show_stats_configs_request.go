package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatsConfigsRequest Request Object
type ShowStatsConfigsRequest struct {

	// - 配置类型 - 目前支持0：热点统计，1：ces上报
	ConfigType int32 `json:"config_type"`
}

func (o ShowStatsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowStatsConfigsRequest", string(data)}, " ")
}
