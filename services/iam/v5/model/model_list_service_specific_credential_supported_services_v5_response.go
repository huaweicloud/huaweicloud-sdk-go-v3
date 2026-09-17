package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSpecificCredentialSupportedServicesV5Response Response Object
type ListServiceSpecificCredentialSupportedServicesV5Response struct {

	// 服务专属凭证所属云服务信息列表。
	Services *[]SupportedService `json:"services,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServiceSpecificCredentialSupportedServicesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSpecificCredentialSupportedServicesV5Response struct{}"
	}

	return strings.Join([]string{"ListServiceSpecificCredentialSupportedServicesV5Response", string(data)}, " ")
}
