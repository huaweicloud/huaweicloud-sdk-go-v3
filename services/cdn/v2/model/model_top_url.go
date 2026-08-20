package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUrl top_url配置
type TopUrl struct {

	// **参数解释：** 配置开关 **约束限制：** 不涉及 **取值范围：** - true：打开 - false：关闭 **默认取值：** 不涉及
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 热点统计配置指标的上报数量 > 如top_url 100、top_url 1000  **约束限制：** 不涉及 **取值范围：** 0-2000 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 热点统计类指标是否支持按状态码上报 **约束限制：** 不涉及 **取值范围：** - true：热点统计类指标支持按状态码上报 - false：热点统计类指标不支持按状态码上报 **默认取值：** 不涉及
	SortByCode *bool `json:"sort_by_code,omitempty"`
}

func (o TopUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrl struct{}"
	}

	return strings.Join([]string{"TopUrl", string(data)}, " ")
}
