package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElasticResourcePoolScaleRecord 扩缩容历史记录接口返回数组中的数据
type ElasticResourcePoolScaleRecord struct {

	// 最大Cu
	MaxCu int32 `json:"max_cu"`

	// 最小CU
	MinCu int32 `json:"min_cu"`

	// 扩缩容后CU
	CurrentCu int32 `json:"current_cu"`

	// 扩缩容前CU
	OriginCu int32 `json:"origin_cu"`

	// 扩缩容预期CU
	TargetCu int32 `json:"target_cu"`

	// 操作完成时间
	RecordTime int64 `json:"record_time"`

	// 扩缩容成功或者失败的状态
	Status string `json:"status"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o ElasticResourcePoolScaleRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElasticResourcePoolScaleRecord struct{}"
	}

	return strings.Join([]string{"ElasticResourcePoolScaleRecord", string(data)}, " ")
}
