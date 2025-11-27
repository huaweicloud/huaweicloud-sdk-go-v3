package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateProvidersResponse Response Object
type ListPrivateProvidersResponse struct {

	// 私有provider的列表。默认以创建时间降序排序。
	Providers *[]PrivateProviderSummary `json:"providers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateProvidersResponse", string(data)}, " ")
}
