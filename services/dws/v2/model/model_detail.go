package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Detail **参数解释**： 规格详细信息。 **取值范围**： 不涉及。
type Detail struct {

	// **参数解释**： 属性类型。 **取值范围**： - vCPU：CPU核心数。 - SATA：普通IO。 - SAS：高IO。 - SSD：超高IO。 - ESSD：极速型SSD。 - GPSSD：通用型SSD。 - LOCAL_DISK：本地盘。 - mem：内存大小。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 属性值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 属性单位。 **取值范围**： 当type为SATA、SAS、SSD、ESSD时，表示磁盘单位GB
	Unit *string `json:"unit,omitempty"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
