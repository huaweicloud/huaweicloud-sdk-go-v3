package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolUpdate
type NodePoolUpdate struct {
	Metadata *NodePoolMetadataUpdate `json:"metadata,omitempty"`

	Spec *NodePoolSpecUpdate `json:"spec,omitempty"`
}

func (o NodePoolUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolUpdate struct{}"
	}

	return strings.Join([]string{"NodePoolUpdate", string(data)}, " ")
}
