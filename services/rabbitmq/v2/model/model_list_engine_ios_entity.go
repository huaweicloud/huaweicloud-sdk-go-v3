package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineIosEntity **参数解释**： 支持的磁盘IO类型信息。 **取值范围**： 不涉及。
type ListEngineIosEntity struct {

	// **参数解释**： 存储类型规格编码。 **取值范围**： - dms.physical.storage.high.v2：高IO类型磁盘。 - dms.physical.storage.ultra.v2：超高IO类型磁盘。 [- dms.physical.storage.general：使用通用型SSD的磁盘类型。](tag:hws,hws_hk,ax) [- dms.physical.storage.extreme：使用极速型SSD的磁盘类型。](tag:hws,hws_hk,ax)
	IoSpec *string `json:"io_spec,omitempty"`

	// **参数解释**： 服务类型。 **取值范围**： evs。
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
