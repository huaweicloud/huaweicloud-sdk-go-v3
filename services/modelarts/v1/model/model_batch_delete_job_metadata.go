package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobMetadata 批量删除训练作业时所需的作业元数据。
type BatchDeleteJobMetadata struct {

	// **参数解释**：训练作业ID，格式为UUID。 **约束限制**：不涉及。 **取值范围**：32位字母、数字与中划线的组合。 **默认取值**：不涉及。
	Id string `json:"id"`
}

func (o BatchDeleteJobMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobMetadata struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobMetadata", string(data)}, " ")
}
