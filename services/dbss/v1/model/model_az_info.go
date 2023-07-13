package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzInfo struct {

	// 可用区名称
	ZoneName string `json:"zone_name"`

	// 可用区编号
	ZoneNumber int32 `json:"zone_number"`

	// 可用区类型
	AzType string `json:"az_type"`

	// 可用区别名
	Alias string `json:"alias"`

	// 可用区别名英文
	AliasUs string `json:"alias_us"`
}

func (o AzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfo struct{}"
	}

	return strings.Join([]string{"AzInfo", string(data)}, " ")
}
