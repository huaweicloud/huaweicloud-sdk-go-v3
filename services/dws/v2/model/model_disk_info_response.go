package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskInfoResponse **参数解释**： 磁盘容量信息对象。 **取值范围**： 不涉及。
type DiskInfoResponse struct {

	// **参数解释**： 已使用百分比。 **取值范围**： 不涉及。
	Percentage *string `json:"percentage,omitempty"`

	// **参数解释**： 节点ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 节点名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 磁盘规格。 **取值范围**： 不涉及。
	DiskCapacity *string `json:"disk_capacity,omitempty"`

	// **参数解释**： 已使用量。 **取值范围**： 不涉及。
	DiskUsed *string `json:"disk_used,omitempty"`
}

func (o DiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskInfoResponse struct{}"
	}

	return strings.Join([]string{"DiskInfoResponse", string(data)}, " ")
}
