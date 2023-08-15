package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRocketMqMigrationTaskRequest Request Object
type DeleteRocketMqMigrationTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *MetadataDeleteReq `json:"body,omitempty"`
}

func (o DeleteRocketMqMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRocketMqMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteRocketMqMigrationTaskRequest", string(data)}, " ")
}
