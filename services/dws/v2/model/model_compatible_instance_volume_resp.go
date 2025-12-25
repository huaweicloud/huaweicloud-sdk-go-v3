package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompatibleInstanceVolumeResp **参数解释**： 容量相关信息。 **取值范围**： 不涉及。
type CompatibleInstanceVolumeResp struct {

	// **参数解释**： 磁盘类型。 **取值范围**： - SATA：普通IO - SAS：高IO - SSD：超高IO - ESSD：极速型SSD - GPSSD：通用型SSD
	Type *string `json:"type,omitempty"`

	// **参数解释**： 已使用空间。 **取值范围**： 不涉及。
	Used *float32 `json:"used,omitempty"`

	// **参数解释**： 总大小。 **取值范围**： 不涉及。
	Size *int32 `json:"size,omitempty"`
}

func (o CompatibleInstanceVolumeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleInstanceVolumeResp struct{}"
	}

	return strings.Join([]string{"CompatibleInstanceVolumeResp", string(data)}, " ")
}
