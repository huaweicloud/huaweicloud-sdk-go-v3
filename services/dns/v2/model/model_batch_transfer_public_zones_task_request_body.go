package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTransferPublicZonesTaskRequestBody struct {

	// **参数解释：** 域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneNames []string `json:"zone_names"`

	// **参数解释：** 对方账号ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TargetUser string `json:"target_user"`
}

func (o BatchTransferPublicZonesTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferPublicZonesTaskRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTransferPublicZonesTaskRequestBody", string(data)}, " ")
}
