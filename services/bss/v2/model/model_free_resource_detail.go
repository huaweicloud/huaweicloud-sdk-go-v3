package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type FreeResourceDetail struct {

	// 资源项ID，一个资源包中会含有多个资源项，一个使用量类型对应一个资源项。
	FreeResourceId *string `json:"free_resource_id,omitempty"`

	// 资源项类型名称。
	FreeResourceTypeName *string `json:"free_resource_type_name,omitempty"`

	// 重置周期，只有quota_reuse_mode为可重置，该字段才有意义。 1：小时2：天3：周4：月5：年
	QuotaReuseCycle *int32 `json:"quota_reuse_cycle,omitempty"`

	// 重置周期类别，只有quota_reuse_mode为可重置，该字段才有意义。 1：按自然周期重置是指重置周期是按照自然月/年来重置，例如如果周期是月，按自然周期重置，表示每个月的1号重置。 2：按订购周期重置。是指重置周期是按照订购时间来重置，例如如果周期是月，按订购周期重置，15号订购了该套餐，表示每个月的15号重置。
	QuotaReuseCycleType *int32 `json:"quota_reuse_cycle_type,omitempty"`

	// 使用量类型名称。
	UsageTypeName *string `json:"usage_type_name,omitempty"`

	// 开始时间，格式UTC。 如果quota_reuse_mode为可重置，则此时间为当前时间所在的重置周期的开始时间。如果quota_reuse_mode为不可重置，则此时间为订购实例的生效时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式UTC。 如果quota_reuse_mode为可重置，则此时间为当前时间所在的重置周期的结束时间。如果quota_reuse_mode为不可重置，则此时间为订购实例的失效时间。
	EndTime *string `json:"end_time,omitempty"`

	// 资源剩余额度，针对可重置资源包，是指当前重置周期内的剩余量。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 资源原始额度，针对可重置资源包，是指每个重置周期内的总量。
	OriginalAmount *decimal.Decimal `json:"original_amount,omitempty"`

	// 度量单位，免费资源套餐额度度量单位。您可以调用查询度量单位列表接口获取。
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o FreeResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreeResourceDetail struct{}"
	}

	return strings.Join([]string{"FreeResourceDetail", string(data)}, " ")
}
