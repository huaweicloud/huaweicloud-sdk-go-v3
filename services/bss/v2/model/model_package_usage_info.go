package model

import (
	"encoding/json"

	"strings"
)

type PackageUsageInfo struct {
	// 订购资源包产品后，系统生成的ID，是这个资源包列表的标识字段。

	OrderInstanceId *string `json:"order_instance_id,omitempty"`
	// 资源类型名称。

	ResourceTypeName *string `json:"resource_type_name,omitempty"`
	// 使用模式。 1：可重置表示购买的按需资源包能够按照一定的周期恢复使用量。例如购买一个1年的按需资源包，使用量是40G，可重置，重置周期为1个月，表示1年内每个月会给予40G的使用量。 2：不可重置。表示购买的按需套餐包的使用量不会恢复。例如购买一个1年的按需资源包，使用量是40G，不可重置，表示1年内一共给予40G的使用量。

	QuotaReuseMode *int32 `json:"quota_reuse_mode,omitempty"`
	// 重置周期，只有quota_reuse_mode为可重置，该字段才有意义。 1：小时2：天3：周4：月5：年

	QuotaReuseCycle *int32 `json:"quota_reuse_cycle,omitempty"`
	// 重置周期类别，只有quota_reuse_mode为可重置，该字段才有意义。 1：按自然周期重置是指重置周期是按照自然月/年来重置，例如如果周期是月，按自然周期重置，表示每个月的1号重置。 2：按订购周期重置。是指重置周期是按照订购时间来重置，例如如果周期是月，按订购周期重置，15号订购了该套餐，表示每个月的15号重置。

	QuotaReuseCycleType *int32 `json:"quota_reuse_cycle_type,omitempty"`
	// 开始时间，格式UTC。 如果quota_reuse_mode为可重置，则此时间为当前时间所在的重置周期的开始时间。如果quota_reuse_mode为不可重置，则此时间为订购实例的生效时间。

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间，格式UTC。 如果quota_reuse_mode为可重置，则此时间为当前时间所在的重置周期的结束时间。如果quota_reuse_mode为不可重置，则此时间为订购实例的失效时间。

	EndTime *string `json:"end_time,omitempty"`
	// 套餐包内资源剩余量。

	Balance float32 `json:"balance,omitempty"`
	// 套餐包的资源总量。

	Total float32 `json:"total,omitempty"`
	// 套餐包资源的度量单位名称。

	MeasurementName *string `json:"measurement_name,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	RegionCode *string `json:"region_code,omitempty"`
}

func (o PackageUsageInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PackageUsageInfo struct{}"
	}

	return strings.Join([]string{"PackageUsageInfo", string(data)}, " ")
}
