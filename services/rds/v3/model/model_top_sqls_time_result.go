package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSqlsTimeResult struct {

	// 对查询计算的二进制哈希值，用于标识具有类似逻辑的查询。
	Id *string `json:"id,omitempty"`

	// 数据类型。取值范围： AvgWorkerTime 平均CPU开销 AvgDuration 平均执行耗时 TotalWorkerTime 总CPU开销 TotalDuration 总执行耗时
	DataType *string `json:"data_type,omitempty"`

	// 耗时时间，单位ms。
	Value *string `json:"value,omitempty"`
}

func (o TopSqlsTimeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlsTimeResult struct{}"
	}

	return strings.Join([]string{"TopSqlsTimeResult", string(data)}, " ")
}
