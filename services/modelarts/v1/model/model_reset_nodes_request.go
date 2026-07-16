package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetNodesRequest struct {

	// **参数解释**：需要重置的节点名称列表。 **约束限制**：不涉及。
	NodeNames []string `json:"nodeNames"`

	RollingConfig *ResetNodesRequestRollingConfig `json:"rollingConfig"`

	NodeConfig *ResetNodesRequestNodeConfig `json:"nodeConfig,omitempty"`
}

func (o ResetNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetNodesRequest struct{}"
	}

	return strings.Join([]string{"ResetNodesRequest", string(data)}, " ")
}
