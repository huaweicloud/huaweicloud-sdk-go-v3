package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAssociateIssueResponseResult **参数解释**： 被取消的关联关系记录详情,包含关联关系的所有属性信息。
type CancelAssociateIssueResponseResult struct {

	// **参数解释**： 关联关系唯一标识。 **取值范围**： 32个字符,由小写字母和数字组成。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释**： 源工作项ID。 **取值范围**： 不涉及。
	IssueId *int32 `json:"issueId,omitempty"`

	// **参数解释**： 源项目数字ID。 **取值范围**： 不涉及。
	ProjectId *int32 `json:"projectId,omitempty"`

	// **参数解释**： 关联类型。 **取值范围**： - associate：关联工作项。
	AssociateType *string `json:"associateType,omitempty"`

	// **参数解释**： 被关联工作项ID。 **取值范围**： 不涉及。
	AssociateIssueId *int32 `json:"associateIssueId,omitempty"`

	// **参数解释**： 被关联项目数字ID。 **取值范围**： 不涉及。
	AssociateProjectId *int32 `json:"associateProjectId,omitempty"`

	// **参数解释**： 关联关系创建时间。 **取值范围**： 格式为yyyy-MM-dd HH:mm:ss。
	CreatedOn *sdktime.SdkTime `json:"createdOn,omitempty"`

	// **参数解释**： 创建该关联关系的用户ID。 **取值范围**： 不涉及。
	AuthorId *int32 `json:"authorId,omitempty"`

	// **参数解释**： 关联关系有效标识。 **取值范围**： - true：关联有效。 - false：关联已失效。
	Flag *bool `json:"flag,omitempty"`
}

func (o CancelAssociateIssueResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAssociateIssueResponseResult struct{}"
	}

	return strings.Join([]string{"CancelAssociateIssueResponseResult", string(data)}, " ")
}
