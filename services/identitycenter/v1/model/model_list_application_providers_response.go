package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationProvidersResponse Response Object
type ListApplicationProvidersResponse struct {

	// 应用程序提供商列表
	ApplicationProviders *[]ApplicationProviderDto `json:"application_providers,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApplicationProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationProvidersResponse", string(data)}, " ")
}
