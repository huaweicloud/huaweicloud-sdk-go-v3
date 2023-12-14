package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitVhostMetadata RocketMQ元数据迁移，RabbitMQ vhost元数据。
type MigrationRabbitVhostMetadata struct {

	// vhost名称。
	Name *string `json:"name,omitempty"`
}

func (o MigrationRabbitVhostMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitVhostMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitVhostMetadata", string(data)}, " ")
}
