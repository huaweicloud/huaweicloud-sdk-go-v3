package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstancesRequest Request Object
type ListResourceInstancesRequest struct {

	// 定值为resource_instances。
	ResourceInstances string `json:"resource_instances"`

	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}
