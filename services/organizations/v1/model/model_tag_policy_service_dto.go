package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// A quota of organization.
type TagPolicyServiceDto struct {

	// The service name of the service.
	ServiceName string `json:"service_name"`

	ResourceTypes []string `json:"resource_types"`

	// resource_type是否支持全量选择，即*
	SupportAll bool `json:"support_all"`
}

func (o TagPolicyServiceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagPolicyServiceDto struct{}"
	}

	return strings.Join([]string{"TagPolicyServiceDto", string(data)}, " ")
}
