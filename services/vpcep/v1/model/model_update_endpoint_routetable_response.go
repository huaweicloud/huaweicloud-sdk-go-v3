package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointRoutetableResponse Response Object
type UpdateEndpointRoutetableResponse struct {

	// 路由表ID列表。 若未指定，返回默认VPC下路由表ID。 更新Gateway类型终端节点服务的终端节点时，显示此参数。
	Routetables *[]string `json:"routetables,omitempty"`

	// 当修改终端节点子网路由表失败时，返回错误提示信息
	Error          *[]RoutetableInfoError `json:"error,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateEndpointRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableResponse", string(data)}, " ")
}
