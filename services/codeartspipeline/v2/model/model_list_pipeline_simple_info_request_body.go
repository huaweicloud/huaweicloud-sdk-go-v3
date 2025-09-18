package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineSimpleInfoRequestBody **参数解释**： 查询流水线信息对象 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ListPipelineSimpleInfoRequestBody struct {

	// **参数解释**： 流水线名字。查询时进行模糊匹配。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PipelineName *string `json:"pipeline_name,omitempty"`

	// **参数解释**： 项目id，有多个值时用逗号分隔，id个数取值[0,10]，非必选。如果该参数有值，则获取对应项目下的流水线列表；如果没有值，则获取用户有权限的所有项目的流水线列表 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectIds *string `json:"project_ids,omitempty"`

	// **参数解释**： 创建人id，有多个值时用逗号分隔，id个数取值[0,10]，非必选 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreatorIds *string `json:"creator_ids,omitempty"`

	// **参数解释**： 执行人id。有多个值时用逗号分隔，id个数取值[0,10]，非必选。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ExecutorIds *string `json:"executor_ids,omitempty"`

	// **参数解释**： 流水线运行状态。 **约束限制**： 不涉及。 **取值范围**： - waiting：等待中。 - running：运行中。 - verifying：待审核。 - suspending：挂起。 - completed：执行完成。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 流水线执行结果。 **约束限制**： 不涉及。 **取值范围**： - success：成功。 - error：失败。 - aborted：终止。 **默认取值**： 不涉及。
	Outcome *string `json:"outcome,omitempty"`

	// **参数解释**： 用于排序的字段，非必选。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**： 排序方式。 **约束限制**： 不涉及。 **取值范围**： - asc：按排序字段升序。 - desc：按排序字段降序。 **默认取值**： 不涉及。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 代码仓地址。仅支持codehub仓库，如：git@codehub.XXX.git **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释**： 偏移量。表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： offset大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **约束限制**： 不涉及。 **取值范围**： 取值[10-50]。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPipelineSimpleInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineSimpleInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ListPipelineSimpleInfoRequestBody", string(data)}, " ")
}
