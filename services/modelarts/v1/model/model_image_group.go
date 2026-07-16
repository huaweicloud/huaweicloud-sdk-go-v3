package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageGroup struct {

	// **参数解释**：镜像名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像创建的时间，单位：UTC毫秒。 **取值范围**：不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释**：镜像所属的SWR组织。 **取值范围**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：镜像最后更新的时间，单位：UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释**：镜像版本个数。 **取值范围**：不涉及。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释**：镜像描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：镜像指导。 **取值范围**：不涉及。
	ReadMe *string `json:"read_me,omitempty"`

	// **参数解释**：镜像图标名称。 **取值范围**：不涉及。
	IconName *string `json:"icon_name,omitempty"`

	// **参数解释**：镜像id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：SWR企业版镜像仓库名称。 **取值范围**：不涉及。
	SwrInstanceName *string `json:"swr_instance_name,omitempty"`

	// **参数解释**：SWR企业版镜像仓库ID。 **取值范围**：不涉及。
	SwrInstanceId *string `json:"swr_instance_id,omitempty"`
}

func (o ImageGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageGroup struct{}"
	}

	return strings.Join([]string{"ImageGroup", string(data)}, " ")
}
