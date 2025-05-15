package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalRegionId 本端 Region ID。
type LocalRegionId struct {

	// RegionID。
	LocalRegionId string `json:"local_region_id"`
}

func (o LocalRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalRegionId struct{}"
	}

	return strings.Join([]string{"LocalRegionId", string(data)}, " ")
}
