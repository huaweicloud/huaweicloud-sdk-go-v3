package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourceUnderNodeRequest struct {

	// 云服务资源；同rms服务的provider
	RfResourceType string `json:"rf_resource_type"`

	// 云服务资源类型；同rms服务的type
	Type string `json:"type"`

	Body *PageResourceListParam `json:"body,omitempty"`
}

func (o ListResourceUnderNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUnderNodeRequest struct{}"
	}

	return strings.Join([]string{"ListResourceUnderNodeRequest", string(data)}, " ")
}
