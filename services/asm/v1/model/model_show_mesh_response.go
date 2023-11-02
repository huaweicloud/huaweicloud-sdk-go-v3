package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMeshResponse Response Object
type ShowMeshResponse struct {

	// API版本，固定值“v1”，该值不可修改
	ApiVersion *string `json:"apiVersion,omitempty"`

	// API类型，固定值“Mesh”或“mesh”，该值不可修改
	Kind *string `json:"kind,omitempty"`

	Metadata *MeshMetadata `json:"metadata,omitempty"`

	Spec *MeshSpec `json:"spec,omitempty"`

	Status         *MeshStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowMeshResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeshResponse struct{}"
	}

	return strings.Join([]string{"ShowMeshResponse", string(data)}, " ")
}
