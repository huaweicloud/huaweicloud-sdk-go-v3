package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 此参数在请求实体中，采用json字符串格式
type QueryLtsStructLogParamsNew struct {

	// sql语句字符串。
	Query string `json:"query"`

	// 查询结果格式。当前仅支持：\"k-v\"。
	Format string `json:"format"`

	TimeRange *TimeRange `json:"time_range"`

	// 返回数据格式，是否为行数据，默认为false。
	WhetherToRows *bool `json:"whether_to_rows,omitempty"`
}

func (o QueryLtsStructLogParamsNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLtsStructLogParamsNew struct{}"
	}

	return strings.Join([]string{"QueryLtsStructLogParamsNew", string(data)}, " ")
}
