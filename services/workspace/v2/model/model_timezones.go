package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Timezones 返回时区配置信息
type Timezones struct {

	// 时区描述
	TimeZoneDesc *string `json:"time_zone_desc,omitempty"`

	// 时区偏移量
	TimeZone *string `json:"time_zone,omitempty"`

	// 时区地名
	TimeZoneName *string `json:"time_zone_name,omitempty"`

	// 时区英文描述
	TimeZoneDescUs *string `json:"time_zone_desc_us,omitempty"`

	// 时区中文描述
	TimeZoneDescCn *string `json:"time_zone_desc_cn,omitempty"`
}

func (o Timezones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Timezones struct{}"
	}

	return strings.Join([]string{"Timezones", string(data)}, " ")
}
