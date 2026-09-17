package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectVo 项目信息。
type ProjectVo struct {

	// **参数解释：** 项目uuid **取值范围：** 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释：** 项目名称 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 项目数字id **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 项目类型 **取值范围：** scrum。
	ProjectType *string `json:"project_type,omitempty"`
}

func (o ProjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectVo struct{}"
	}

	return strings.Join([]string{"ProjectVo", string(data)}, " ")
}
