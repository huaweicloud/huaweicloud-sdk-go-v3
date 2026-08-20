package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainCountryStatResponse Response Object
type ShowDomainCountryStatResponse struct {

	// **参数解释：** 查询数据类型 > 汇总或明细数据  **取值范围：** - summary：查询汇总数据 - detail：查询数据详情
	Action *string `json:"action,omitempty"`

	// **参数解释：** 查询起始时间 **取值范围：** 相对于UTC 1970-01-01到当前时间相隔的毫秒数
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 查询结束时间 **取值范围：** 相对于UTC 1970-01-01到当前时间相隔的毫秒数
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 统计指标类型 **取值范围：** - flux：流量 - req_num：请求总数
	StatType *string `json:"stat_type,omitempty"`

	// **参数解释：** 按指定的分组方式组织的数据 **取值范围：** - domain：按域名分组 - country：按国际&地区分组 - province：按省份分组 - isp：按运营商分组 **默认取值：** 默认不分组
	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainCountryStatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainCountryStatResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainCountryStatResponse", string(data)}, " ")
}
