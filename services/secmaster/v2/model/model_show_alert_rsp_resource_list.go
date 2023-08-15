package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlertRspResourceList struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The name, display only
	Type *string `json:"type,omitempty"`

	// Id value
	DomainId *string `json:"domain_id,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	RegionId *string `json:"region_id,omitempty"`

	// Id value
	EpId *string `json:"ep_id,omitempty"`

	// The name, display only
	EpName *string `json:"ep_name,omitempty"`

	// Id value
	Tags *string `json:"tags,omitempty"`
}

func (o ShowAlertRspResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspResourceList struct{}"
	}

	return strings.Join([]string{"ShowAlertRspResourceList", string(data)}, " ")
}
