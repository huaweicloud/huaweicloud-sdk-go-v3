package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KillOpRule struct {

	// killOp规则ID。
	Id *string `json:"id,omitempty"`

	// Sql语句操作类型。 最多支持同时选择6种语句类型。同时选择多种类型时，匹配任意一种类型时规则生效。 - insert，表示插入语句。  - update，表示更新语句。  - query，表示查询语句。  - command，表示命令语句。  - remove，表示删除语句。  - getmore，表示获取更多数据语句。
	OperationTypes *string `json:"operation_types,omitempty"`

	// killOp规则状态。  - ENABLED，规则生效中。 - DISABLED，规则禁用中。
	Status *string `json:"status,omitempty"`

	// 表命名空间。取值格式：库名.表名。同时配置多组信息时，匹配任意一组信息时规则生效 - 目前仅支持配置一组信息 - 可为空，表示不做限制。 - 单独库名，表示对某个库下的所有集合生效。 - 库名.表名，表示对具体库下的具体的集合生效。
	Namespaces *string `json:"namespaces,omitempty"`

	// 客户端连接IP。只支持IPV4。可为空，表示不做限制。最多支持配置5个IP。同时配置多个IP时，匹配任意一个IP时规则生效。
	ClientIps *string `json:"client_ips,omitempty"`

	// 执行计划。 默认值空，表示不做限制。  - COLLSCAN。 - SORT_KEY_GENERATOR。 - SKIP。 - LIMIT。 - GEO_NEAR_2DSPHERE。 - GEO_NEAR_2D。 - AGGREGATE。 - OR。
	PlanSummary *string `json:"plan_summary,omitempty"`

	// 最大并发数。 取值： 不能为负数，可为空，默认为0，表示不做限制， 最小值为1， 最大值为100000。secs_running和max_concurrency不可同时为0。
	MaxConcurrency *int32 `json:"max_concurrency,omitempty"`

	// 单条操作最大运行时长。取值： 不能为负数，可为空，默认为0，表示不做限制。单位：s。最小值为2， 最大值为86400。secs_running和max_concurrency不可同时为0。
	SecsRunning *int32 `json:"secs_running,omitempty"`

	// 节点类型。  - mongos_shard，表示同时在mongos和shard节点生效。 - mongos，表示只在集群mongos节点生效。 - shard，表示只在集群shard节点生效。 - replica，表示只在副本集节点生效。
	NodeType *string `json:"node_type,omitempty"`
}

func (o KillOpRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillOpRule struct{}"
	}

	return strings.Join([]string{"KillOpRule", string(data)}, " ")
}
