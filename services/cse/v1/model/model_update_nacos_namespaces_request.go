package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNacosNamespacesRequest Request Object
type UpdateNacosNamespacesRequest struct {

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 命名空间ID
	Namespace string `json:"namespace"`

	// 命名空间名，支持非@、#、$、%、^、&、*，不超过128个字符。
	NamespaceShowName string `json:"namespace_show_name"`

	// 命名空间描述，不超过256个字符。
	NamespaceDesc string `json:"namespace_desc"`
}

func (o UpdateNacosNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNacosNamespacesRequest struct{}"
	}

	return strings.Join([]string{"UpdateNacosNamespacesRequest", string(data)}, " ")
}
