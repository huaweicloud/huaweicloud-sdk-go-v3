package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectProtectedTagPo struct {

	// **参数解释：** 保护分支唯一标识。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 来源。 **取值范围：** 字符串长度不少于1，不超过1000。
	Source *string `json:"source,omitempty"`

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 保护分支名称 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 事件列表。
	Actions *[]ProjectProtectedActionResultDto `json:"actions,omitempty"`
}

func (o ProjectProtectedTagPo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectProtectedTagPo struct{}"
	}

	return strings.Join([]string{"ProjectProtectedTagPo", string(data)}, " ")
}
