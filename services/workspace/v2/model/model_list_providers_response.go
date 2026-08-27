package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProvidersResponse Response Object
type ListProvidersResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 供应商配置列表项。
	Items          *[]ProviderInfoVo `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListProvidersResponse", string(data)}, " ")
}
