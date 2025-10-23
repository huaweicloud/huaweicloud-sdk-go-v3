package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BccOrganizationPolicy BccOrganizationPolicy
type BccOrganizationPolicy struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	Detail *BccOrganizationPolicyDetail `json:"detail,omitempty"`
}

func (o BccOrganizationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BccOrganizationPolicy struct{}"
	}

	return strings.Join([]string{"BccOrganizationPolicy", string(data)}, " ")
}
