package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorInfo 规格内容信息。
type FlavorInfo struct {

	// 参数解释：并发数。单位：个。
	Connection int32 `json:"connection"`

	// 参数解释：新建数。单位：个。
	Cps int32 `json:"cps"`

	// 参数解释：每秒查询数。单位：个。  约束限制：仅7层LB有该指标。
	Qps *int32 `json:"qps,omitempty"`

	// 参数解释：带宽。单位：Kbit/s。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 参数解释：当前flavor对应的lcu数量。单位：个。  > LCU是用来衡量独享型ELB处理性能综合指标，LCU值越大，性能越好。
	Lcu *int32 `json:"lcu,omitempty"`

	// 参数解释：https新建连接数。单位：个。  约束限制：仅7层LB有该指标。
	HttpsCps *int32 `json:"https_cps,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
