package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceGroupsServicesResourcesResponse Response Object
type ListResourceGroupsServicesResourcesResponse struct {

	// 资源总数
	Count *int32 `json:"count,omitempty"`

	// 资源分组资源列表
	Resources      *[]GetResourceGroupResources `json:"resources,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListResourceGroupsServicesResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupsServicesResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupsServicesResourcesResponse", string(data)}, " ")
}
