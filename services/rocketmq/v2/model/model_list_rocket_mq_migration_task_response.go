package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRocketMqMigrationTaskResponse Response Object
type ListRocketMqMigrationTaskResponse struct {

	// **参数解释**： 元数据迁移任务总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 元数据迁移任务列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Task           *[]MetadataTask `json:"task,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListRocketMqMigrationTaskResponse", string(data)}, " ")
}
