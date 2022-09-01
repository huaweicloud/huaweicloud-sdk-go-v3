package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 攻击来源区域统计
type GeoItem struct {

	// 攻击来源区域
	Key *string `json:"key,omitempty" xml:"key"`

	// 数量
	Num *int32 `json:"num,omitempty" xml:"num"`
}

func (o GeoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeoItem struct{}"
	}

	return strings.Join([]string{"GeoItem", string(data)}, " ")
}
