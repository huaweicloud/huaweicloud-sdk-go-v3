package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowFileTreeResultChildren struct {

	// **参数解释**: 文件或文件夹名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 访问地址。 **取值范围**: 不涉及。
	Uri *string `json:"uri,omitempty"`

	// **参数解释**: 路径。 **取值范围**: 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释**: 更新时间，时间格式：yyyy-MM-dd HH:mm:ss。 **取值范围**: 不涉及。
	Modified *string `json:"modified,omitempty"`

	// **参数解释**: 是否为文件夹。 **取值范围**: true：文件夹。 false：文件。
	Folder *bool `json:"folder,omitempty"`

	// **参数解释**: 修改人。 **取值范围**: 不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释**: 是否存在下一层。 **取值范围**: true：是。 false：否。
	HasChild *bool `json:"has_child,omitempty"`
}

func (o ShowFileTreeResultChildren) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileTreeResultChildren struct{}"
	}

	return strings.Join([]string{"ShowFileTreeResultChildren", string(data)}, " ")
}
