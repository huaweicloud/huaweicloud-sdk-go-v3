package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowVo 工作项流程流转请求参数。
type WorkItemFlowVo struct {

	// **参数解释**： 工作项唯一ID。可以通过[查询工作项列表](ListIpdProjectIssues.xml)或者[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**id**字段的值就是工作项ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项类型。 **约束限制**： 不涉及。 **取值范围**： RR、IR、AR、SR、Bug、FE、Task、US、Epic、SF **默认取值**： 不涉及。
	IssueCategory string `json:"issue_category"`

	// **参数解释**： 工作项流转code。可以通过[查询工作项流程信息](ShowIssueWorkItemFlowDetail.xml)接口获取。 响应消息体中的**next_flow**数组为工作流流转线，根据**from_code**当前状态和**to_code**目标状态找到匹配的流转线，流转线的**code**字段的值就是工作项流转code。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FlowCode string `json:"flow_code"`

	// **参数解释**： 工作项唯一Id数组。可以通过[查询工作项列表](ListIpdProjectIssues.xml)或者[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**id**字段的值就是工作项ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IssueIds *[]string `json:"issue_ids,omitempty"`

	// **参数解释**： 流转中配置上下文信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProcessContext map[string]interface{} `json:"process_context,omitempty"`
}

func (o WorkItemFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowVo", string(data)}, " ")
}
