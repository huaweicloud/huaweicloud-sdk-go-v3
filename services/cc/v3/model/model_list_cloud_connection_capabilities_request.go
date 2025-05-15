package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionCapabilitiesRequest Request Object
type ListCloudConnectionCapabilitiesRequest struct {

	// 类型。
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ListCloudConnectionCapabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionCapabilitiesRequest", string(data)}, " ")
}
