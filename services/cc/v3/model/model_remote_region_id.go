package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteRegionId 对端 Region ID。
type RemoteRegionId struct {

	// RegionID。
	RemoteRegionId string `json:"remote_region_id"`
}

func (o RemoteRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteRegionId struct{}"
	}

	return strings.Join([]string{"RemoteRegionId", string(data)}, " ")
}
