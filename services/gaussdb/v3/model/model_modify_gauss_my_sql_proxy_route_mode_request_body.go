package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyGaussMySqlProxyRouteModeRequestBody proxy实例修改路由模式请求体
type ModifyGaussMySqlProxyRouteModeRequestBody struct {

	// 数据库代理路由模式。  取值范围： - 0，表示权重负载模式。 - 1，表示负载均衡模式（数据库主节点不接受读请求）。 - 2，表示负载均衡模式（数据库主节点接受读请求）。
	RouteMode int32 `json:"route_mode"`

	// 主节点权重： - 如果路由模式为0，取值为0~1000。 - 如果路由模式为1，取值为0。 - 如果路由模式为2，取值为1。
	MasterWeight *int32 `json:"master_weight,omitempty"`

	// 只读节点权重配置信息。
	ReadonlyNodes *[]ModifyProxyRouteWeightReadonlyNode `json:"readonly_nodes,omitempty"`
}

func (o ModifyGaussMySqlProxyRouteModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGaussMySqlProxyRouteModeRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyGaussMySqlProxyRouteModeRequestBody", string(data)}, " ")
}
