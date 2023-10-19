package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 云服务类型，固定为'hws.service.type.ccm'
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型，CA为\"hws.resource.type.pca.duration\"
	ResourceType string `json:"resource_type"`

	// 资源规格编码，CA为\"ca.duration\"
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
