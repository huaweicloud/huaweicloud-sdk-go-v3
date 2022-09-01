package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNodeResponse struct {

	// API类型，固定值“Node”，该值不可修改。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	Metadata *NodeMetadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *NodeSpec `json:"spec,omitempty" xml:"spec"`

	Status         *NodeStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeResponse", string(data)}, " ")
}
