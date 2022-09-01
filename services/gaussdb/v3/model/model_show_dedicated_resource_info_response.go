package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDedicatedResourceInfoResponse struct {

	// 专属资源池id
	Id *string `json:"id,omitempty" xml:"id"`

	// 专属资源池名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 引擎名称
	EngineName *string `json:"engine_name,omitempty" xml:"engine_name"`

	// 可用区
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty" xml:"availability_zone_ids"`

	// 资源规格类型
	Architecture *string `json:"architecture,omitempty" xml:"architecture"`

	// 专属资源池状态
	Status *string `json:"status,omitempty" xml:"status"`

	DedicatedComputeInfo *DedicatedComputeInfo `json:"dedicated_compute_info,omitempty" xml:"dedicated_compute_info"`

	DedicatedStorageInfo *DedicatedStorageInfo `json:"dedicated_storage_info,omitempty" xml:"dedicated_storage_info"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowDedicatedResourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedResourceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowDedicatedResourceInfoResponse", string(data)}, " ")
}
