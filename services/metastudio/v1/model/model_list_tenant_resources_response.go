package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantResourcesResponse Response Object
type ListTenantResourcesResponse struct {

	// 资源用量列表
	Resources *[]ResourceItemInfo `json:"resources,omitempty"`

	// 资源总数。
	Count float32 `json:"count,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTenantResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantResourcesResponse", string(data)}, " ")
}
