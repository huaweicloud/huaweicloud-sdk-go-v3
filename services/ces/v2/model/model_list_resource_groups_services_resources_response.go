package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceGroupsServicesResourcesResponse Response Object
type ListResourceGroupsServicesResourcesResponse struct {

	// **参数解释** 资源总数。 **取值范围** 0-10000
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
