package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateEndpointRoutetableResponse struct {
	// 路由表ID列表。节点的白名单。 此参数可以添加IPv4或CIDR： ● 当取值不为空时，表示将白 名单更新为取值所示内容。 ● 当取值为空时，表示删除所 有白名单。 默认为空列表。

	Routetables *[]string `json:"routetables,omitempty"`

	Error          *RoutetableInfoError `json:"error,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateEndpointRoutetableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableResponse", string(data)}, " ")
}
