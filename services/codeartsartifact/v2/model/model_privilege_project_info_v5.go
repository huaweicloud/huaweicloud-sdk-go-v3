package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivilegeProjectInfoV5 struct {

	// **参数解释**：  项目的序号。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 项目的名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 项目的状态（该参数无返回值，请忽略）。 **取值范围**： 该参数无返回值，请忽略。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 项目的创建时间。 **取值范围**： 格式为yyyy-MM-dd HH:mm:ss。
	CreatedTime *string `json:"created_time,omitempty"`

	// **参数解释**： 仓库数量。 **取值范围**： 不涉及。
	RepositoryNum *int32 `json:"repository_num,omitempty"`

	// **参数解释**： 项目总数（该参数无返回值，请忽略）。 **取值范围**： 该参数无返回值，请忽略。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 仓库ID。 **取值范围**： 不涉及。
	RepositoryId *string `json:"repository_id,omitempty"`

	// **参数解释**： 项目所属的区域。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 项目的序号列表（该参数无返回值，请忽略）。 **取值范围**： 该参数无返回值，请忽略。
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释**： 项目总数。 **取值范围**： 不涉及。
	TotalRecords *int32 `json:"total_records,omitempty"`

	// **参数解释**： 请求参数中的仓库ID是否关联到本项目。 **取值范围**： - true：关联到本项目。 - false：未关联到本项目。
	Associated *bool `json:"associated,omitempty"`

	// **参数解释**： 项目的信息。 **取值范围**： 不涉及。
	ProjectInfo *[]PrivilegeProjectInfoV5 `json:"project_info,omitempty"`
}

func (o PrivilegeProjectInfoV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegeProjectInfoV5 struct{}"
	}

	return strings.Join([]string{"PrivilegeProjectInfoV5", string(data)}, " ")
}
