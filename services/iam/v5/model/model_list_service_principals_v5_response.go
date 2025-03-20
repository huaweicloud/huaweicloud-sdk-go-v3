package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicePrincipalsV5Response Response Object
type ListServicePrincipalsV5Response struct {

	// 服务主体列表。
	ServicePrincipals *[]ServicePrincipalMetadata `json:"service_principals,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServicePrincipalsV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePrincipalsV5Response struct{}"
	}

	return strings.Join([]string{"ListServicePrincipalsV5Response", string(data)}, " ")
}
