package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HyperNode 超节点
type HyperNode struct {

	// v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// HyperNode
	Kind *string `json:"kind,omitempty"`

	Metadata *HyperNodeMetadata `json:"metadata,omitempty"`

	Spec *HyperNodeSpec `json:"spec,omitempty"`

	Status *HyperNodeStatus `json:"status,omitempty"`
}

func (o HyperNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNode struct{}"
	}

	return strings.Join([]string{"HyperNode", string(data)}, " ")
}
