package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreeResourceRecord struct {

	// 资源抵扣时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如\"2021-10-13T10:01:49Z\"。
	DeductTime *string `json:"deduct_time,omitempty"`

	// 资源项ID，一个资源包中会含有多个资源项，一个使用量类型对应一个资源项。
	FreeResourceId *string `json:"free_resource_id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型编码。例如ECS的VM为“hws.resource.type.vm”。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 资源标签。
	ResourceTag *string `json:"resource_tag,omitempty"`

	// 产品ID，即资源包ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品名称，即资源包名称。
	ProductName *string `json:"product_name,omitempty"`

	// 使用量类型。
	UsageTypeCode *string `json:"usage_type_code,omitempty"`

	// 资源抵扣前余量。
	AvailableAmount *string `json:"available_amount,omitempty"`

	// 资源抵扣后余量。
	RemainingAmount *string `json:"remaining_amount,omitempty"`

	// 抵扣量。
	UsedAmount *string `json:"used_amount,omitempty"`

	// 度量单位，免费资源套餐额度度量单位。您可以调用查询度量单位列表接口获取。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 资源使用的开始时间，UTC时间。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 资源使用的结束时间，UTC时间。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o FreeResourceRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreeResourceRecord struct{}"
	}

	return strings.Join([]string{"FreeResourceRecord", string(data)}, " ")
}
