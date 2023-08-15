package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterRegion struct {

	// 域间实例的ID。
	Id *string `json:"id,omitempty"`

	// 域间实例本段的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 域间实例本段的RegionID。
	LocalRegionId *string `json:"local_region_id,omitempty"`

	// 域间实例对段的RegionID。
	RemoteRegionId *string `json:"remote_region_id,omitempty"`
}

func (o InterRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterRegion struct{}"
	}

	return strings.Join([]string{"InterRegion", string(data)}, " ")
}
