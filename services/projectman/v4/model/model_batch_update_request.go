package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRequest struct {

	// **参数解释：** 模块的负责人数字id，通过[获取指定项目的成员用户列表](ListProjectMembersV4.xml)接口获取，响应消息体中的**user_num_id**字段的值就是模块的负责人数字id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AssignedToId *string `json:"assigned_to_id,omitempty"`

	// **参数解释：** 工作项id，可通过[高级查询工作项](ListIssuesV4.xml)接口获取，响应消息体中的**id**字段的值就是工作项id。 **约束限制：** 长度在1位到10位之间的纯数字。 **取值范围：** 最小长度：1，最大长度：10。 **默认取值：** 不涉及。
	IssueIds *string `json:"issue_ids,omitempty"`

	// **参数解释**： 项目的32位uuid，项目唯一标识，通过[查询项目列表](ListProjectsV4.xml)接口获取，响应消息体中的**project_id**字段的值就是项目ID。 **约束限制**： 32位的数字和字母组成的字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o BatchUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateRequest", string(data)}, " ")
}
