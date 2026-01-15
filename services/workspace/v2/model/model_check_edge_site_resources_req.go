package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEdgeSiteResourcesReq 校验边缘站点资源请求体。
type CheckEdgeSiteResourcesReq struct {

	// 可用区id。
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 规格id。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 需要的资源数量。
	ResourceCounts *int32 `json:"resource_counts,omitempty"`

	// 磁盘列表。包含系统盘。
	Volumes *[]CheckEdgeSiteResourcesVolume `json:"volumes,omitempty"`
}

func (o CheckEdgeSiteResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEdgeSiteResourcesReq struct{}"
	}

	return strings.Join([]string{"CheckEdgeSiteResourcesReq", string(data)}, " ")
}
