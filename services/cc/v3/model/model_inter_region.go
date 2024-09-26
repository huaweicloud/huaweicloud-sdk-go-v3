package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterRegion struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// 域间实例本端的RegionID。
	LocalRegionId *string `json:"local_region_id,omitempty"`

	// 域间实例对端的RegionID。
	RemoteRegionId *string `json:"remote_region_id,omitempty"`
}

func (o InterRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterRegion struct{}"
	}

	return strings.Join([]string{"InterRegion", string(data)}, " ")
}
