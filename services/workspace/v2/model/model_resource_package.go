package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcePackage 按需资源包。
type ResourcePackage struct {

	// 资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 产品描述<语言，各语言对应的产品名>。
	Description map[string]string `json:"description,omitempty"`

	// 产品状态，normal/onSell：正常、sellout：售空、abandon：下线。
	Status *string `json:"status,omitempty"`

	// 周期类型，MONTH：月。
	PeriodType *string `json:"period_type,omitempty"`

	// 数量，单位:月。
	PeriodValue *int32 `json:"period_value,omitempty"`

	// 资源包单位，hour：小时。
	MeasurementName *string `json:"measurement_name,omitempty"`

	// 资源包总量。
	Threshold *int32 `json:"threshold,omitempty"`
}

func (o ResourcePackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePackage struct{}"
	}

	return strings.Join([]string{"ResourcePackage", string(data)}, " ")
}
