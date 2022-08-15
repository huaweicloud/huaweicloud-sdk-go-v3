package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeCondition struct {

	// |参数名称：时间单位1：天2：月| |参数的约束及描述：时间单位1：天2：月|
	TimeMeasureId int32 `json:"time_measure_id"`

	// |参数名称：查询开始时间，必须是日期格式。当time_measure_id值为1或为空时，格式为YYYY-MM-DD当time_measure_id值为2时，格式为YYYY-MM| |参数约束及描述：查询开始时间，必须是日期格式。当time_measure_id值为1或为空时，格式为YYYY-MM-DD当time_measure_id值为2时，格式为YYYY-MM|
	BeginTime string `json:"begin_time"`

	// |参数名称：查询结束时间：必须是日期格式。当time_measure_id值为1或为空时，格式为YYYY-MM-DD当time_measure_id值为2时，格式为YYYY-MM| |参数约束及描述：查询结束时间：必须是日期格式。当time_measure_id值为1或为空时，格式为YYYY-MM-DD当time_measure_id值为2时，格式为YYYY-MM|
	EndTime string `json:"end_time"`
}

func (o TimeCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeCondition struct{}"
	}

	return strings.Join([]string{"TimeCondition", string(data)}, " ")
}
