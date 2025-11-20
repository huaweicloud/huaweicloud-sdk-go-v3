package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobMetadata
type JobMetadata struct {

	// **参数解释**： 任务的ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 任务的创建时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**： 任务的更新时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o JobMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMetadata struct{}"
	}

	return strings.Join([]string{"JobMetadata", string(data)}, " ")
}
