package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreeResourceV3 struct {

	// 资源项ID，一个资源包中会含有多个资源项，一个使用量类型对应一个资源项。
	FreeResourceId *string `json:"free_resource_id,omitempty"`

	// 使用量类型名称。
	UsageTypeName *string `json:"usage_type_name,omitempty"`

	// 资源剩余额度，针对可重置资源包，是指当前重置周期内的剩余量。
	Amount *string `json:"amount,omitempty"`

	// 资源原始额度，针对可重置资源包，是指每个重置周期内的总量。
	OriginalAmount *string `json:"original_amount,omitempty"`

	// 度量单位，免费资源套餐额度度量单位。您可以调用查询度量单位列表接口获取。
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o FreeResourceV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreeResourceV3 struct{}"
	}

	return strings.Join([]string{"FreeResourceV3", string(data)}, " ")
}
