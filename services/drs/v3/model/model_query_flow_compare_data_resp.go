package model

import (
	"encoding/json"

	"strings"
)

// 灾备初始化对象详情信息
type QueryFlowCompareDataResp struct {
	// 任务总数

	TotalRecord *int64 `json:"total_record,omitempty"`
	// 数据生成时间

	CreateTime *string `json:"create_time,omitempty"`
	// 对比结果

	List *[]StructDetailVo `json:"list,omitempty"`
}

func (o QueryFlowCompareDataResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryFlowCompareDataResp struct{}"
	}

	return strings.Join([]string{"QueryFlowCompareDataResp", string(data)}, " ")
}
