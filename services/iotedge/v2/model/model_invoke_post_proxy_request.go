package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokePostProxyRequest Request Object
type InvokePostProxyRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 第三方应用IA ID
	IaId string `json:"ia_id"`

	// 第三方IA服务资源地址
	IaUri string `json:"ia_uri"`

	Body *interface{} `json:"body,omitempty"`
}

func (o InvokePostProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokePostProxyRequest struct{}"
	}

	return strings.Join([]string{"InvokePostProxyRequest", string(data)}, " ")
}
