package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllProvidersResponse Response Object
type ListAllProvidersResponse struct {

	// 云服务详情列表
	ResourceProviders *[]ResourceProviderResponse `json:"resource_providers,omitempty"`

	// 当前支持的云服务总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListAllProvidersResponse", string(data)}, " ")
}
