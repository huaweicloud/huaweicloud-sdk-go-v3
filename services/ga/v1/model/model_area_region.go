package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 区域。
type AreaRegion struct {

	// 区域所属地区，取值： - OUTOFCM： 中国大陆以外 - CM：中国大陆
	Area *string `json:"area,omitempty"`

	// 区域ID列表。
	Regions *[]string `json:"regions,omitempty"`
}

func (o AreaRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaRegion struct{}"
	}

	return strings.Join([]string{"AreaRegion", string(data)}, " ")
}
