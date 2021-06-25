package model

import (
	"encoding/json"

	"strings"
)

// 此参数在请求实体中，采用json字符串格式
type QueryLtsStructLogParamsNew struct {
	// sql语句字符串。

	Query string `json:"query"`
	// 查询结果格式。当前仅支持：\"k-v\"。

	Format string `json:"format"`
	// 时间范围信息。

	TimeRange *interface{} `json:"time_range"`
}

func (o QueryLtsStructLogParamsNew) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryLtsStructLogParamsNew struct{}"
	}

	return strings.Join([]string{"QueryLtsStructLogParamsNew", string(data)}, " ")
}
