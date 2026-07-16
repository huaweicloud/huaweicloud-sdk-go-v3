package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchActionDevServerIds 批量操作Lite Server实例ID列表。
type BatchActionDevServerIds struct {

	// **参数解释**：Lite Server实例ID。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及
	Id string `json:"id"`
}

func (o BatchActionDevServerIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchActionDevServerIds struct{}"
	}

	return strings.Join([]string{"BatchActionDevServerIds", string(data)}, " ")
}
