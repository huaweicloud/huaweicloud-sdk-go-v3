package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokePatchProxyRequest Request Object
type InvokePatchProxyRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 第三方应用IA ID
	IaId string `json:"ia_id"`

	// 第三方IA服务资源地址
	IaUri string `json:"ia_uri"`

	Body *interface{} `json:"body,omitempty"`
}

func (o InvokePatchProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokePatchProxyRequest struct{}"
	}

	return strings.Join([]string{"InvokePatchProxyRequest", string(data)}, " ")
}
