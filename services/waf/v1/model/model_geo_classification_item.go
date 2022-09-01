package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeoClassificationItem struct {

	// GeoItem的总数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// GeoItem详细信息
	Items *[]GeoItem `json:"items,omitempty" xml:"items"`
}

func (o GeoClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeoClassificationItem struct{}"
	}

	return strings.Join([]string{"GeoClassificationItem", string(data)}, " ")
}
