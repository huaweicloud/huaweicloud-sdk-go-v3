package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionDto 局点信息
type RegionDto struct {

	// 局点ID
	RegionId string `json:"region_id"`
}

func (o RegionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionDto struct{}"
	}

	return strings.Join([]string{"RegionDto", string(data)}, " ")
}
