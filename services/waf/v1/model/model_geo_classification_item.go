package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeoClassificationItem struct {

	// GeoItem的总数量
	Total *int32 `json:"total,omitempty"`

	// GeoItem详细信息
	Items *[]GeoItem `json:"items,omitempty"`
}

func (o GeoClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeoClassificationItem struct{}"
	}

	return strings.Join([]string{"GeoClassificationItem", string(data)}, " ")
}
