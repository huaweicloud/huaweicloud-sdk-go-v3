package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthRef **参数解释**：共享带宽的ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
type BandwidthRef struct {

	// **参数解释**：共享带宽的ID。
	Id string `json:"id"`
}

func (o BandwidthRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthRef struct{}"
	}

	return strings.Join([]string{"BandwidthRef", string(data)}, " ")
}
