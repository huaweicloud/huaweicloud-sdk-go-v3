package model

import (
	"encoding/json"

	"strings"
)

//
type QueryCompareResultReq struct {
	// 任务id。

	JobId string `json:"job_id"`
	// 请求查询结果的对象级对比任务id。

	ObjectLevelCompareId *string `json:"object_level_compare_id,omitempty"`
	// 请求查询结果的行对比任务id。

	LineCompareId *string `json:"line_compare_id,omitempty"`
	// 请求查询结果的内容对比任务id。

	ContentCompareId *string `json:"content_compare_id,omitempty"`
	// 分页查询的当前页码，对查询对比任务的结果生效。

	CurrentPage int32 `json:"current_page"`
	// 分页查询的每页个数，对查询对比任务的结果生效。

	PerPage int32 `json:"per_page"`
}

func (o QueryCompareResultReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryCompareResultReq struct{}"
	}

	return strings.Join([]string{"QueryCompareResultReq", string(data)}, " ")
}
