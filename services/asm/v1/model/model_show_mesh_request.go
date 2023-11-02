package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMeshRequest Request Object
type ShowMeshRequest struct {

	// 网格ID
	MeshId string `json:"mesh_id"`

	// 网格所属ProjectID
	XApplyProjectID *string `json:"X-Apply-ProjectID,omitempty"`
}

func (o ShowMeshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeshRequest struct{}"
	}

	return strings.Join([]string{"ShowMeshRequest", string(data)}, " ")
}
