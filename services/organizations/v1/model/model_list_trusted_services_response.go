package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustedServicesResponse Response Object
type ListTrustedServicesResponse struct {

	// 启用与组织集成的服务主体列表。
	TrustedServices *[]TrustedServiceDto `json:"trusted_services,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTrustedServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedServicesResponse struct{}"
	}

	return strings.Join([]string{"ListTrustedServicesResponse", string(data)}, " ")
}
