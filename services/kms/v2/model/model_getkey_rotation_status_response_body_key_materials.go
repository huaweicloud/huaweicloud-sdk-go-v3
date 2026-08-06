package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetkeyRotationStatusResponseBodyKeyMaterials struct {

	// **参数解释：** 密钥材料ID **取值范围：** uuid格式
	MaterialId *string `json:"material_id,omitempty"`

	// **参数解释：** 计费ID **取值范围：** 不涉及
	ChargeId *string `json:"charge_id,omitempty"`

	// **参数解释：** 密钥材料创建时间 **取值范围：** 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** 密钥材料过期时间 **取值范围：** 不涉及
	ExpirationTime *string `json:"expiration_time,omitempty"`

	// **参数解释：** 密钥材料状态 **取值范围：** 0：等待轮转状态；2：启用状态
	State *int32 `json:"state,omitempty"`
}

func (o GetkeyRotationStatusResponseBodyKeyMaterials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetkeyRotationStatusResponseBodyKeyMaterials struct{}"
	}

	return strings.Join([]string{"GetkeyRotationStatusResponseBodyKeyMaterials", string(data)}, " ")
}
