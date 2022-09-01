package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaResource struct {

	// 资源类型
	QuotaKey *string `json:"quotaKey,omitempty" xml:"quotaKey"`

	// 配额值
	QuotaLimit *int32 `json:"quotaLimit,omitempty" xml:"quotaLimit"`

	// 已创建的资源个数
	Used *int32 `json:"used,omitempty" xml:"used"`

	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// 局点ID。若资源不涉及此参数，则不返回该参数。
	RegionId *string `json:"regionId,omitempty" xml:"regionId"`

	// 可用区ID。若资源不涉及此参数，则不返回该参数。
	AvailabilityZoneId *string `json:"availabilityZoneId,omitempty" xml:"availabilityZoneId"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
