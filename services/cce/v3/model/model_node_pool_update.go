package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NodePoolUpdate struct {
	Metadata *NodePoolMetadataUpdate `json:"metadata" xml:"metadata"`

	Spec *NodePoolSpecUpdate `json:"spec" xml:"spec"`
}

func (o NodePoolUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolUpdate struct{}"
	}

	return strings.Join([]string{"NodePoolUpdate", string(data)}, " ")
}
