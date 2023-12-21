package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BandwidthTypeOption struct {

	// 共享带宽类型名称
	Name *string `json:"name,omitempty"`

	// 共享带宽类型
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// 站点ID
	SiteId *string `json:"site_id,omitempty"`
}

func (o BandwidthTypeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthTypeOption struct{}"
	}

	return strings.Join([]string{"BandwidthTypeOption", string(data)}, " ")
}
