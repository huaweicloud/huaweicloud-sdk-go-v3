package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstallCmdRequest Request Object
type CreateInstallCmdRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 节点架构
	Arch string `json:"arch"`

	Body *CreateInstallCmdRequestDto `json:"body,omitempty"`
}

func (o CreateInstallCmdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstallCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateInstallCmdRequest", string(data)}, " ")
}
