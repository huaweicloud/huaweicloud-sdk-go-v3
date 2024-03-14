package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedServicesResponse Response Object
type ListDelegatedServicesResponse struct {

	// 账号是其委托管理员的服务。
	DelegatedServices *[]DelegatedServiceDto `json:"delegated_services,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDelegatedServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDelegatedServicesResponse struct{}"
	}

	return strings.Join([]string{"ListDelegatedServicesResponse", string(data)}, " ")
}
