package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteModuleRequestBody 批量删除工作项请求体。
type BatchDeleteModuleRequestBody struct {

	// **参数解释**： 项目的32位uuid，项目唯一标识，通过[查询项目列表](ListProjectsV4.xml)接口获取，响应消息体中的**project_id**字段的值就是项目ID。 **约束限制**： 32位的数字和字母组成的字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 工作项id，可通过[高级查询工作项](ListIssuesV4.xml)接口获取，响应消息体中的**id**字段的值就是工作项id。 **约束限制：** 长度在1位到10位之间的纯数字。 **取值范围：** 最小长度：1，最大长度：10。 **默认取值：** 不涉及。
	IssueIds *string `json:"issue_ids,omitempty"`
}

func (o BatchDeleteModuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteModuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteModuleRequestBody", string(data)}, " ")
}
