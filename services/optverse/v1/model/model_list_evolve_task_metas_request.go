package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEvolveTaskMetasRequest Request Object
type ListEvolveTaskMetasRequest struct {

	// **参数解释**： 排序规则，目前默认创建时间降序。 **约束限制**： 不涉及 **取值范围**： - DESC：降序 - ASC：升序 **默认取值**： DESC
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**： 关联的算法设计项目。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AlgorithmId string `json:"algorithm_id"`

	// **参数解释**： 任务名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 任务状态列表。 **约束限制**： 不涉及 **取值范围**： - DRAFT：草稿 - PENDING：初始化 - RUNNING：运行中 - FINISHED：已完成 - STOPPED：已停止 - FAILED：失败 **默认取值**： 不涉及
	StatusList *[]string `json:"status_list,omitempty"`

	// **参数解释**： 用户名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。 **约束限制**： 不涉及 **取值范围**： [1,1000] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。 **约束限制**： 不涉及 **取值范围**： [0,100000000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListEvolveTaskMetasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvolveTaskMetasRequest struct{}"
	}

	return strings.Join([]string{"ListEvolveTaskMetasRequest", string(data)}, " ")
}
