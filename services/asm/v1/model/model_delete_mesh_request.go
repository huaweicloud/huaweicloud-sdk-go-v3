package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMeshRequest Request Object
type DeleteMeshRequest struct {

	// 网格ID。
	MeshId string `json:"mesh_id"`
}

func (o DeleteMeshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMeshRequest struct{}"
	}

	return strings.Join([]string{"DeleteMeshRequest", string(data)}, " ")
}
