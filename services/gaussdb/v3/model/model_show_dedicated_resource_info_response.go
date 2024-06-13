package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDedicatedResourceInfoResponse Response Object
type ShowDedicatedResourceInfoResponse struct {

	// 专属资源池ID。
	Id *string `json:"id,omitempty"`

	// 专属资源池名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 可用区。
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`

	// 资源规格类型。
	Architecture *string `json:"architecture,omitempty"`

	// 专属资源池状态。
	Status *string `json:"status,omitempty"`

	DedicatedComputeInfo *DedicatedComputeInfo `json:"dedicated_compute_info,omitempty"`

	DedicatedStorageInfo *DedicatedStorageInfo `json:"dedicated_storage_info,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowDedicatedResourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedResourceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowDedicatedResourceInfoResponse", string(data)}, " ")
}
