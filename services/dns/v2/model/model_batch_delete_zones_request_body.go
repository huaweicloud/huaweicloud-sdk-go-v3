package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteZonesRequestBody struct {

	// **参数解释：** 待删除域名的类型。 **约束限制：** 不涉及。 **取值范围：** - public：公网域名 - private：内网域名  **默认取值：** 不涉及。
	ZoneType string `json:"zone_type"`

	// **参数解释：** 待删除域名ID列表。 **约束限制：** 最多支持50个。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneIds []string `json:"zone_ids"`
}

func (o BatchDeleteZonesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteZonesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteZonesRequestBody", string(data)}, " ")
}
