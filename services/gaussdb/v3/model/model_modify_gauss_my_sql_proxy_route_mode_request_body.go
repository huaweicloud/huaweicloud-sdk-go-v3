package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 是否开启新增节点自动加入该Proxy。如果需要设置是否开启新增节点自动加入该Proxy，请联系客服人员添加白名单，加入白名单后，方可输入该字段。  取值范围： - ON：开启。 - OFF：关闭。
	NewNodeAutoAddStatus *string `json:"new_node_auto_add_status,omitempty"`

	// 新增节点的读权重：    - 如果路由模式为0，新增节点自动加入为ON，取值为0~1000。 - 如果路由模式不为0或新增节点自动加入为OFF，则可不输入读权重。 - 如果路由模式不为0或新增节点自动加入为OFF，则可不输入读权重。
	NewNodeWeight *int32 `json:"new_node_weight,omitempty"`
}

func (o ModifyGaussMySqlProxyRouteModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGaussMySqlProxyRouteModeRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyGaussMySqlProxyRouteModeRequestBody", string(data)}, " ")
}
