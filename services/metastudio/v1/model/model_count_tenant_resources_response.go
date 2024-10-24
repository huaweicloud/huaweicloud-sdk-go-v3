package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTenantResourcesResponse Response Object
type CountTenantResourcesResponse struct {

	// 资源总数列表
	ResourcesCount *[]ResourcesCount `json:"resources_count,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountTenantResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTenantResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountTenantResourcesResponse", string(data)}, " ")
}
