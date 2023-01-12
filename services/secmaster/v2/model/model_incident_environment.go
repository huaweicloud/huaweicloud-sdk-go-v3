package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境信息
type IncidentEnvironment struct {

	// The name, display only
	VendorType *string `json:"vendor_type,omitempty"`

	// Id value
	DomainId *string `json:"domain_id,omitempty"`

	// Id value
	RegionId *string `json:"region_id,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`
}

func (o IncidentEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentEnvironment struct{}"
	}

	return strings.Join([]string{"IncidentEnvironment", string(data)}, " ")
}
