package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReadableNodeInfos 可读节点信息。
type ReadableNodeInfos struct {

	// 可读节点IP。
	DataIp *string `json:"data_ip,omitempty"`

	// 可读节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 可读节点名称。
	NodeName *string `json:"node_name,omitempty"`
}

func (o ReadableNodeInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadableNodeInfos struct{}"
	}

	return strings.Join([]string{"ReadableNodeInfos", string(data)}, " ")
}
