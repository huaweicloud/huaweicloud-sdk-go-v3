package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AzInfo 可用区信息。
type AzInfo struct {

	// 可用区名称。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	I18n *I18n `json:"i18n,omitempty"`

	// 是否为默认可用分区。
	DefaultAvailabilityZone *bool `json:"default_availability_zone,omitempty"`
}

func (o AzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfo struct{}"
	}

	return strings.Join([]string{"AzInfo", string(data)}, " ")
}
