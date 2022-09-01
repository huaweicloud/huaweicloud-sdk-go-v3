package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateNodePoolResponse struct {

	// API类型，固定值“NodePool”。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v3”。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	Metadata *NodePoolMetadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *NodePoolSpec `json:"spec,omitempty" xml:"spec"`

	Status         *NodePoolStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolResponse struct{}"
	}

	return strings.Join([]string{"CreateNodePoolResponse", string(data)}, " ")
}
