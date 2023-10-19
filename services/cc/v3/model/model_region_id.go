package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionId Region ID。
type RegionId struct {

	// RegionID。
	RegionId string `json:"region_id"`
}

func (o RegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionId struct{}"
	}

	return strings.Join([]string{"RegionId", string(data)}, " ")
}
