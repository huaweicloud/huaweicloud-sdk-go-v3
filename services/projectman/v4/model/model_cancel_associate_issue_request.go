package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAssociateIssueRequest 取消关联工作项请求对象
type CancelAssociateIssueRequest struct {

	// **参数解释**： 源项目UUID。标识执行取消关联操作的源工作项所属项目。 **约束限制**： 32位UUID字符串,必填字段。 **取值范围**： 32个字符,由小写字母和数字组成。 **默认取值**： 不涉及。
	ProjectUUId string `json:"projectUUId"`

	// **参数解释**： 目标项目UUID。标识被取消关联工作项所属的项目;跨项目取消时必填。 **约束限制**： 32位UUID字符串。 **取值范围**： 32个字符,由小写字母和数字组成。 **默认取值**： 不涉及。
	AttachProjectUUId *string `json:"attachProjectUUId,omitempty"`

	// **参数解释**： 源工作项ID。即需要解除关联关系的工作项唯一ID。 **约束限制**： 工作项必须存在且未被归档。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IssueId int32 `json:"issueId"`

	// **参数解释**： 待取消关联的目标工作项ID。 **约束限制**： 必须与源工作项已存在关联关系;不存在则返回错误码DEV_21_400806。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AttachIssueId int32 `json:"attachIssueId"`
}

func (o CancelAssociateIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAssociateIssueRequest struct{}"
	}

	return strings.Join([]string{"CancelAssociateIssueRequest", string(data)}, " ")
}
