package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionViewVoV5 struct {

	// **参数解释**: id。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 项目id。 **取值范围**: 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 文件id。 **取值范围**: 不涉及。
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**: 文件名。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 文件路径。 **取值范围**: 不涉及。
	Path *string `json:"path,omitempty"`
}

func (o VersionViewVoV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionViewVoV5 struct{}"
	}

	return strings.Join([]string{"VersionViewVoV5", string(data)}, " ")
}
