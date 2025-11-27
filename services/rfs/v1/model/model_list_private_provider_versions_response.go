package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateProviderVersionsResponse Response Object
type ListPrivateProviderVersionsResponse struct {

	// 私有provider版本的列表。默认以创建时间降序排序。
	Versions *[]PrivateProviderVersionSummary `json:"versions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateProviderVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateProviderVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateProviderVersionsResponse", string(data)}, " ")
}
