package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListMetricDataRequestBody
type BatchListMetricDataRequestBody struct {

	// **参数解释** 指标数据 **约束限制** 包含的指标数据对象个数为[1,500]
	Metrics []MetricInfo `json:"metrics"`

	Period *BatchPeriod `json:"period"`

	Filter *Filter `json:"filter"`

	// **参数解释** 查询数据起始时间，UNIX时间戳，单位毫秒 **约束限制** 当period为1时，若(to- from) >4*3600*1000，则from调整为 to - 4*3600*1000 当period为300时，若(to - from) >24*3600*1000，则from调整为 to - 24*3600*1000 当period为1200时，若(to - from) >3*24*3600*1000，则from调整为 to - 3*24*3600*1000 当period为3600时，若(to -from) > 10*24*3600*1000，则from调整为 to -10*24*3600*1000 当period为14400时，若(to - from) >30*24*3600*1000，则from调整为 to - 30*24*3600*1000 当period为86400时，若(to -from) > 180*24*3600*1000，则from调整为 to - 180*24*3600*1000 **取值范围** 毫秒级时间戳范围为[1111111111111,9999999999999] **默认取值** 不涉及
	From int64 `json:"from"`

	// **参数解释** 查询数据截止时间，UNIX时间戳，单位毫秒 **约束限制** from 必须小于to **取值范围** 毫秒级时间戳范围为[1111111111111,9999999999999] **默认取值** 不涉及
	To int64 `json:"to"`
}

func (o BatchListMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequestBody", string(data)}, " ")
}
