package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVariableGroupDetailResponse Response Object
type ShowVariableGroupDetailResponse struct {

	// **参数解释**： 参数组名称。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 32位字符，由数字和字母组成。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 参数组名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 参数组描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 参数列表。 **取值范围**： 不涉及。
	Variables *[]QueryVariableGroupDetailRespVariables `json:"variables,omitempty"`

	// **参数解释**： 关联的流水线。 **取值范围**： 不涉及。
	RelatedPipelines *[]QueryVariableGroupDetailRespRelatedPipelines `json:"related_pipelines,omitempty"`

	// **参数解释**： 创建人ID。 **取值范围**： 32位字符，由数字和字母组成。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 编辑人ID。 **取值范围**： 32位字符，由数字和字母组成。
	UpdaterId *string `json:"updater_id,omitempty"`

	// **参数解释**： 创建人名称。 **取值范围**： 不涉及。
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**： 编辑人名称。 **取值范围**： 不涉及。
	UpdaterName *string `json:"updater_name,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *int32 `json:"create_time,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdateTime     *int32 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVariableGroupDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVariableGroupDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowVariableGroupDetailResponse", string(data)}, " ")
}
