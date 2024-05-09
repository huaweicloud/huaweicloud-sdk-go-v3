package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ZoneDetail struct {

	// 地区编码，如CN。不区分大小写，统一转为大写处理
	Code *string `json:"code,omitempty"`

	// 地区名称
	Name *string `json:"name,omitempty"`
}

func (o ZoneDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZoneDetail struct{}"
	}

	return strings.Join([]string{"ZoneDetail", string(data)}, " ")
}
