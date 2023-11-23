package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillOpRuleRuleListRequest Request Object
type ShowKillOpRuleRuleListRequest struct {

	// 实例ID，可以调用“[查询实例列表和详情](x-wc://file=zh-cn_topic_0000001369935045.xml)”接口获取。如果未申请实例，可以调用“[创建实例](x-wc://file=zh-cn_topic_0000001369734929.xml)”接口创建。
	InstanceId string `json:"instance_id"`

	// Sql语句操作类型。  - insert，表示插入语句。  - update，表示更新语句。  - query，表示查询语句。  - command，表示命令语句。  - remove，表示删除语句。  - getmore，表示获取更多数据语句。
	OperationTypes *string `json:"operation_types,omitempty"`

	// 表命名空间。取值格式：库名.表名。 - 可为空，表示不做限制。 - 单独库名，表示对某个库下的所有集合生效。 - 库名.表名，表示对具体库下的具体的集合生效。
	Namespaces *string `json:"namespaces,omitempty"`

	// killOp规则状态。  - ENABLED，规则生效中。 - DISABLED，规则禁用中。
	Status *string `json:"status,omitempty"`

	// 执行计划。 默认值空，表示不做限制。  - COLLSCAN。 - SORT_KEY_GENERATOR。 - SKIP。 - LIMIT。 - GEO_NEAR_2DSPHERE。 - GEO_NEAR_2D。 - AGGREGATE。 - OR。
	PlanSummary *string `json:"plan_summary,omitempty"`

	// 索引位置，偏移量。  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。 取值必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。 - 取值范围: 1~100。 - 不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowKillOpRuleRuleListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillOpRuleRuleListRequest struct{}"
	}

	return strings.Join([]string{"ShowKillOpRuleRuleListRequest", string(data)}, " ")
}
