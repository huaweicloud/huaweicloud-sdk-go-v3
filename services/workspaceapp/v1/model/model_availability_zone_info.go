package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailabilityZoneInfo struct {

	// 可用分区编码。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 可用分区名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 可用分区国际化信息。
	I18n map[string]string `json:"i18n,omitempty"`

	SoldOut *SoldOutInfo `json:"sold_out,omitempty"`

	// 指定当前分区下自定义支持的产品ID列表，如果为空则支持所有套餐。
	ProductIds *[]string `json:"product_ids,omitempty"`

	// 是否可见。
	Visible *bool `json:"visible,omitempty"`

	// 是否默认可用分区。
	DefaultAvailabilityZone *bool `json:"default_availability_zone,omitempty"`
}

func (o AvailabilityZoneInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZoneInfo struct{}"
	}

	return strings.Join([]string{"AvailabilityZoneInfo", string(data)}, " ")
}
