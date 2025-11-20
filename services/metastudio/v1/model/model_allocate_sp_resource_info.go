package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllocateSpResourceInfo 分配给租户的资源。
type AllocateSpResourceInfo struct {

	// 资源分配id。
	AssignmentRecordId *string `json:"assignment_record_id,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源计费类型。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源数量。
	ResourceNum *float64 `json:"resource_num,omitempty"`

	// 资源已使用数量。
	ResourceUsedNum *float64 `json:"resource_used_num,omitempty"`

	// 资源分配时间（UTC时间）,unix时间,精确到秒。
	ResourceAllocateTime *int64 `json:"resource_allocate_time,omitempty"`

	// 资源到期时间(UTC时间),unix时间,精确到秒。
	ResourceExpireTime *int64 `json:"resource_expire_time,omitempty"`
}

func (o AllocateSpResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllocateSpResourceInfo struct{}"
	}

	return strings.Join([]string{"AllocateSpResourceInfo", string(data)}, " ")
}
