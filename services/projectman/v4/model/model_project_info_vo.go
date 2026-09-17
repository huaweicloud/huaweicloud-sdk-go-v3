package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectInfoVo 项目详情对象
type ProjectInfoVo struct {

	// **参数解释**： 项目ID。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 项目名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 项目类型。 **取值范围**： - ipd：IPD项目 - scrum：scrum项目 - xboard：看板项目
	ProjectType *string `json:"project_type,omitempty"`

	// **参数解释**： 项目空间ID。 **取值范围**： 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： IPD项目模型Id。 **取值范围**： 10001（系统设备类） 10002（独立软件类） 10003（云服务类型）
	ModelId *string `json:"model_id,omitempty"`

	// **参数解释**： 该项目是否接受外部RR（原始需求）。 **取值范围**： - 0：不接受外部RR - 1：接受外部RR
	AcceptRr *int32 `json:"accept_rr,omitempty"`

	// **参数解释**： 项目类型，用于区分项目和项目群。 **取值范围**： - Project：项目 - Group：项目群
	Category *string `json:"category,omitempty"`

	// **参数解释**： 项目创建人名称。 **取值范围**： 不涉及。
	CreatedByName *string `json:"created_by_name,omitempty"`
}

func (o ProjectInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectInfoVo struct{}"
	}

	return strings.Join([]string{"ProjectInfoVo", string(data)}, " ")
}
