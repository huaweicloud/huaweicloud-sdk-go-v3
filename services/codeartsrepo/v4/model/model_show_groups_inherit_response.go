package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupsInheritResponse Response Object
type ShowGroupsInheritResponse struct {

	// **参数解释：** 代码组id。 **取值范围：** 字符串长度不少于1，不超过1000。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 资源类型设置。 **取值范围：** 字符串长度不少于1，不超过1000。
	SourceSetting *string `json:"source_setting,omitempty"`

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 向上继承是否可编辑。
	UpwardInheritEditable *bool `json:"upward_inherit_editable,omitempty"`
	HttpStatusCode        int   `json:"-"`
}

func (o ShowGroupsInheritResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsInheritResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupsInheritResponse", string(data)}, " ")
}
