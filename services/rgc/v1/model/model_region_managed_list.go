package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionManagedList 区域纳管情况。
type RegionManagedList struct {

	// 区域名字。
	Region *string `json:"region,omitempty"`

	// available or unavailable。
	RegionStatus *string `json:"region_status,omitempty"`
}

func (o RegionManagedList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionManagedList struct{}"
	}

	return strings.Join([]string{"RegionManagedList", string(data)}, " ")
}
