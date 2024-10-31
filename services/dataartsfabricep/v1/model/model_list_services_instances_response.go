package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicesInstancesResponse Response Object
type ListServicesInstancesResponse struct {

	// 符合条件的service Istance总数
	Total *int32 `json:"total,omitempty"`

	// 符合条件的service Instance列表
	ServiceInstances *[]ServiceInstanceBriefInfo `json:"service_instances,omitempty"`
	HttpStatusCode   int                         `json:"-"`
}

func (o ListServicesInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListServicesInstancesResponse", string(data)}, " ")
}
