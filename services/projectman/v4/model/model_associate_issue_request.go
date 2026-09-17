package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIssueRequest 关联工作项请求对象
type AssociateIssueRequest struct {

	// **参数解释**： 源项目UUID。标识执行关联操作的源工作项所属项目。 **约束限制**： 32位UUID字符串,必填字段。 **取值范围**： 32个字符,由小写字母和数字组成。 **默认取值**： 不涉及。
	ProjectUUId string `json:"projectUUId"`

	// **参数解释**： 目标项目UUID。标识待关联工作项所属的项目;跨项目关联时必填,同项目关联时可省略。 **约束限制**： 32位UUID字符串;若与projectUUId不同则视为跨项目关联。 **取值范围**： 32个字符,由小写字母和数字组成。 **默认取值**： 不涉及。
	AttachProjectUUId *string `json:"attachProjectUUId,omitempty"`

	// **参数解释**： 源工作项ID。即需要建立关联关系的工作项唯一ID。 **约束限制**： 工作项必须存在且未被归档。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IssueId int32 `json:"issueId"`

	// **参数解释**： 待关联工作项ID列表。本次操作需要新增关联关系的目标工作项ID集合。 **约束限制**： 每个元素为字符串形式的工作项ID(服务端自动转换为整数);不能包含issueId自身;单工作项关联总数受系统上限约束。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AssociatedIssueIdList *[]string `json:"associatedIssueIdList,omitempty"`

	// **参数解释**： 待取消关联工作项ID列表。本次操作需要解除关联关系的目标工作项ID集合;可在同一次请求中混合使用以支持关联关系调整。 **约束限制**： 每个元素为字符串形式的工作项ID;仅处理已存在的关联关系。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UnassociatedIssueIdList *[]string `json:"unassociatedIssueIdList,omitempty"`
}

func (o AssociateIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIssueRequest struct{}"
	}

	return strings.Join([]string{"AssociateIssueRequest", string(data)}, " ")
}
