package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DsConfig struct {

	// 数据选择(天)。
	TimeInterval *int32 `json:"time_interval,omitempty" xml:"time_interval"`

	// 物品类别。
	CategoryType *string `json:"category_type,omitempty" xml:"category_type"`
}

func (o DsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DsConfig struct{}"
	}

	return strings.Join([]string{"DsConfig", string(data)}, " ")
}
