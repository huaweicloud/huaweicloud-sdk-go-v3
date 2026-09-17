package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueFlowRequest struct {

	// **参数解释**： 状态id。 **约束限制**： 不涉及。 **取值范围**： 1（新建） 2（进行中） 3（已解决） 4（测试中） 5（已关闭） 6（已拒绝）。 **默认取值**： 不涉及。
	StatusId *int32 `json:"status_id,omitempty"`

	// **参数解释：** 模块的负责人id，通过[获取指定项目的成员用户列表](ListProjectMembersV4.xml)接口获取，响应消息体中的**user_id**字段的值就是模块的负责人id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AssignedToId *string `json:"assigned_to_id,omitempty"`

	// **参数解释：** 与日志记录相关的备注或注释。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Notes *string `json:"notes,omitempty"`

	// **参数解释**： 项目的32位uuid，项目唯一标识，通过[查询项目列表](ListProjectsV4.xml)接口获取，响应消息体中的**project_id**字段的值就是项目ID。 **约束限制**： 32位的数字和字母组成的字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectUUId *string `json:"projectUUId,omitempty"`

	// **参数解释：** 工作项id，可通过[高级查询工作项](ListIssuesV4.xml)接口获取，响应消息体中的**id**字段的值就是工作项id。 **约束限制：** 长度在1位到10位之间的纯数字。 **取值范围：** 最小长度：1，最大长度：10。 **默认取值：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 项目状态。 **约束限制**： 不涉及。 **取值范围**： scrum。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o IssueFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueFlowRequest struct{}"
	}

	return strings.Join([]string{"IssueFlowRequest", string(data)}, " ")
}
