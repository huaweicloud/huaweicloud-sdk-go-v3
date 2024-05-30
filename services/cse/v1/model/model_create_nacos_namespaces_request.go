package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNacosNamespacesRequest Request Object
type CreateNacosNamespacesRequest struct {

	// 微服务引擎的实例ID
	XEngineId string `json:"x-engine-id"`

	// 企业项目ID
	XEnterpriseProjectID string `json:"X-Enterprise-Project-ID"`

	// 命名空间ID，仅支持大小写字母、数字、短划线（-）和下划线（_），不超过128个字符。
	CustomNamespaceId string `json:"custom_namespace_id"`

	// 命名空间名，支持非@、#、$、%、^、&、*，不超过128个字符。
	NamespaceName string `json:"namespace_name"`

	// 命名空间描述，不超过256个字符。
	NamespaceDesc *string `json:"namespace_desc,omitempty"`
}

func (o CreateNacosNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNacosNamespacesRequest struct{}"
	}

	return strings.Join([]string{"CreateNacosNamespacesRequest", string(data)}, " ")
}
