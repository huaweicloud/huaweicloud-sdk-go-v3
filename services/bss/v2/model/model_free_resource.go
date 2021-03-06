package model

import (
	"encoding/json"

	"strings"
)

type FreeResource struct {
	// 资源项ID。

	FreeResourceId *string `json:"free_resource_id,omitempty"`
	// 使用量类型名称。

	UsageTypeName *string `json:"usage_type_name,omitempty"`
	// 资源剩余额度，针对可重置套餐包，是指当前重置周期内的剩余量。

	Amount *float64 `json:"amount,omitempty"`
	// 资源原始额度，针对可重置套餐包，是指每个重置周期内的总量。

	OriginalAmount *float64 `json:"original_amount,omitempty"`
	// 度量单位，免费资源套餐额度度量单位。您可以调用查询度量单位列表接口获取。

	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o FreeResource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FreeResource struct{}"
	}

	return strings.Join([]string{"FreeResource", string(data)}, " ")
}
