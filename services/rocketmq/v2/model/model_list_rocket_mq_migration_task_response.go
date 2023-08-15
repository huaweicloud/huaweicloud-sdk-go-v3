package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRocketMqMigrationTaskResponse Response Object
type ListRocketMqMigrationTaskResponse struct {

	// 元数据迁移任务总数。
	Total *int32 `json:"total,omitempty"`

	// 元数据迁移任务列表。
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
