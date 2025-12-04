package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineIosEntity **参数解释**： 支持的磁盘IO类型信息。 **取值范围**： 不涉及。
type ListEngineIosEntity struct {

	// **参数解释**： 磁盘IO编码。 **取值范围**： 不涉及。
	IoSpec *string `json:"io_spec,omitempty"`

	// **参数解释**： 磁盘类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 可用区。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 不可用区。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty"`
}

func (o ListEngineIosEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineIosEntity struct{}"
	}

	return strings.Join([]string{"ListEngineIosEntity", string(data)}, " ")
}
