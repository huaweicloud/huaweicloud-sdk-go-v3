package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegisteredServicesForAuthSchemaV5Response Response Object
type ListRegisteredServicesForAuthSchemaV5Response struct {

	// 服务名称缩写列表。
	ServiceCodes *[]string `json:"service_codes,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRegisteredServicesForAuthSchemaV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegisteredServicesForAuthSchemaV5Response struct{}"
	}

	return strings.Join([]string{"ListRegisteredServicesForAuthSchemaV5Response", string(data)}, " ")
}
