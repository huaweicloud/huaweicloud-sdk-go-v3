package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNacosNamespacesRequest Request Object
type DeleteNacosNamespacesRequest struct {

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 命名空间ID
	NamespaceId string `json:"namespace_id"`
}

func (o DeleteNacosNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNacosNamespacesRequest struct{}"
	}

	return strings.Join([]string{"DeleteNacosNamespacesRequest", string(data)}, " ")
}
