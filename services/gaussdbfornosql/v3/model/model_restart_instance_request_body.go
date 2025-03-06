package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestartInstanceRequestBody struct {

	// 待重启的节点ID。 仅GeminiDB Redis云原生部署模式集群实例支持传入节点ID重启对应节点。 不传则重启整个实例。
	NodeId *string `json:"node_id,omitempty"`
}

func (o RestartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequestBody", string(data)}, " ")
}
