package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerScaleDownRequest 超节点缩容请求
type ServerScaleDownRequest struct {

	// 超节点ID
	Id *string `json:"id,omitempty"`

	// 规格信息
	Flavor *string `json:"flavor,omitempty"`

	// 缩容节点id
	ServerIds *[]string `json:"server_ids,omitempty"`

	// 资源规格信息
	ResourceFlavor *string `json:"resource_flavor,omitempty"`
}

func (o ServerScaleDownRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerScaleDownRequest struct{}"
	}

	return strings.Join([]string{"ServerScaleDownRequest", string(data)}, " ")
}
