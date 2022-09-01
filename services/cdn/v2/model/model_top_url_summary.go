package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// top url 详情数据
type TopUrlSummary struct {

	// URL名称。
	Url *string `json:"url,omitempty" xml:"url"`

	// 对应查询类型的值。（流量单位：Byte）
	Value *int64 `json:"value,omitempty" xml:"value"`

	// 查询起始时间戳。
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 查询结束时间戳
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`

	// 参数类型支持：flux(流量)，req_num(请求总数)。
	StatType *string `json:"stat_type,omitempty" xml:"stat_type"`
}

func (o TopUrlSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrlSummary struct{}"
	}

	return strings.Join([]string{"TopUrlSummary", string(data)}, " ")
}
