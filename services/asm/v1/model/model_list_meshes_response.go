package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMeshesResponse Response Object
type ListMeshesResponse struct {

	// API版本，固定值“v1”，该值不可修改
	ApiVersion *string `json:"apiVersion,omitempty"`

	// API类型，固定值“MeshList”，该值不可修改
	Kind *string `json:"kind,omitempty"`

	// 网格列表
	Items          *[]Mesh `json:"items,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMeshesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeshesResponse struct{}"
	}

	return strings.Join([]string{"ListMeshesResponse", string(data)}, " ")
}
