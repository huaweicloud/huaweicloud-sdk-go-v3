package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeGetProxyRequest Request Object
type InvokeGetProxyRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 第三方应用IA ID
	IaId string `json:"ia_id"`

	// 第三方IA服务资源地址
	IaUri string `json:"ia_uri"`

	Body *interface{} `json:"body,omitempty"`
}

func (o InvokeGetProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeGetProxyRequest struct{}"
	}

	return strings.Join([]string{"InvokeGetProxyRequest", string(data)}, " ")
}
