package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableResourceRequest Request Object
type ShowAvailableResourceRequest struct {
	AvailabilityZone []string `json:"availability_zone"`

	FlavorId *[]string `json:"flavor_id,omitempty"`

	DecProjectId *[]string `json:"dec_project_id,omitempty"`

	CheckLimit *[]string `json:"check_limit,omitempty"`

	Expectation *[]string `json:"expectation,omitempty"`

	ResourceType *[]string `json:"resource_type,omitempty"`
}

func (o ShowAvailableResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableResourceRequest", string(data)}, " ")
}
