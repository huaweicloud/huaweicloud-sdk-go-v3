package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReinstallCmdRequest Request Object
type CreateReinstallCmdRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 是否启用TPM
	EnableTpm *bool `json:"enable_tpm,omitempty"`

	Body *CreateReinstallCmdRequestBody `json:"body,omitempty"`
}

func (o CreateReinstallCmdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReinstallCmdRequest struct{}"
	}

	return strings.Join([]string{"CreateReinstallCmdRequest", string(data)}, " ")
}
