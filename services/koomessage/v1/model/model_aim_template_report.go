package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimTemplateReport 查询模板返回体。
type AimTemplateReport struct {

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 统计开始时间。样例为：1970-01-01T00:00:00Z。
	StartTime *string `json:"start_time,omitempty"`

	// 实际已解析数。
	ResolvingTimes *int32 `json:"resolving_times,omitempty"`

	// 统计结束时间。样例为：1970-01-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 消息曝光数。
	ExposeUv *int64 `json:"expose_uv,omitempty"`

	// 消息曝光次数。
	ExposePv *int64 `json:"expose_pv,omitempty"`

	// 消息点击数。
	ClickUv *int64 `json:"click_uv,omitempty"`

	// 消息点击次数。
	ClickPv *int64 `json:"click_pv,omitempty"`
}

func (o AimTemplateReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimTemplateReport struct{}"
	}

	return strings.Join([]string{"AimTemplateReport", string(data)}, " ")
}
