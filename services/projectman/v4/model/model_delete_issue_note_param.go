package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIssueNoteParam 删除工作项评论请求对象
type DeleteIssueNoteParam struct {

	// **参数解释**： 评论ID。标识需要删除的工作项评论唯一记录。 **约束限制**： 评论必须存在，且当前用户必须是该评论的创建者。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id int32 `json:"id"`

	// **参数解释**： 项目ID。标识当前评论所属的项目，用于权限校验与服务可用性校验。 **约束限制**： 32位UUID字符串，且必须与评论对应工作项所属项目保持一致。 **取值范围**： 32个字符，由小写字母和数字组成。 **默认取值**： 不涉及。
	ProjectId string `json:"projectId"`

	// **参数解释**： 工作项类型。标识当前操作对应的工作项类型分类。 **约束限制**： 不涉及。 **取值范围**： - scrum：Scrum项目类型工作项 - 其他类型取值请参考实际业务定义。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o DeleteIssueNoteParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIssueNoteParam struct{}"
	}

	return strings.Join([]string{"DeleteIssueNoteParam", string(data)}, " ")
}
