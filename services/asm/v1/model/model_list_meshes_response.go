package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMeshesResponse Response Object
type ListMeshesResponse struct {
	ApiVersion *string `json:"apiVersion,omitempty"`

	Kind *string `json:"kind,omitempty"`

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
