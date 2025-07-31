package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateZoneStatusRequestBody struct {

	// **参数解释：** 域名状态。 **约束限制：** 不涉及。 **取值范围：** - ENABLE：启用解析 - DISABLE：暂停解析  **默认取值：** 不涉及。
	Status string `json:"status"`
}

func (o UpdateZoneStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateZoneStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateZoneStatusRequestBody", string(data)}, " ")
}
