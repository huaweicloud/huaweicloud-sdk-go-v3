package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DssPool **参数解释**： 专属分布式存储池详情。 **取值范围**： 不涉及。
type DssPool struct {

	// **参数解释**： 专属分布式存储池名称。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 专属分布式存储池ID。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 专属分布式存储池的存储类型。 **取值范围**： - SSD：超高IO专属分布式存储池。
	Type string `json:"type"`

	// **参数解释**： 项目ID。获取方法请参见[获取项目ID](dws_02_0011.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 专属分布式存储池所属可用区。 **取值范围**： 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**： 申请的专属分布式存储容量，单位TB。 **取值范围**： 不涉及。
	Capacity int32 `json:"capacity"`

	// **参数解释**： 专属分布式存储池的状态。 **取值范围**： - available：专属分布式存储池处于可用状态。 - deploying：专属分布式存储池处于正在部署的过程中，不可使用。 - extending：专属分布式存储池处于正在扩容的过程中，可使用。
	Status string `json:"status"`

	// **参数解释**： 专属分布式存储池的创建时间。 **取值范围**： 时间格式：UTC YYYY-MM-DDTHH:MM:SS
	CreatedAt string `json:"created_at"`
}

func (o DssPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DssPool struct{}"
	}

	return strings.Join([]string{"DssPool", string(data)}, " ")
}
