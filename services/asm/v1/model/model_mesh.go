package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Mesh struct {

	// API版本，固定值“v1”，该值不可修改
	ApiVersion string `json:"apiVersion"`

	// API类型，固定值“Mesh”或“mesh”，该值不可修改
	Kind string `json:"kind"`

	Metadata *MeshMetadata `json:"metadata"`

	Spec *MeshSpec `json:"spec"`

	Status *MeshStatus `json:"status,omitempty"`
}

func (o Mesh) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Mesh struct{}"
	}

	return strings.Join([]string{"Mesh", string(data)}, " ")
}
