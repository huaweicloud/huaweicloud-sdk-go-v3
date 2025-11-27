package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatsConfigsRequest Request Object
type ShowStatsConfigsRequest struct {

	// - 配置类型 - 目前支持0：热点统计，1：ces上报
	ConfigType int32 `json:"config_type"`

	// **参数解释：** 分页查询每页的数量 **约束限制：** 不涉及 **取值范围：** 1-1000 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询偏移量，表示跳过多少个数据开始查询 **约束限制：** 不涉及 **取值范围：** 0-65535 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowStatsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowStatsConfigsRequest", string(data)}, " ")
}
