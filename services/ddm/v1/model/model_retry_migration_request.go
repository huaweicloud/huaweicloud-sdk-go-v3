package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryMigrationRequest Request Object
type RetryMigrationRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	// 任务流id
	JobId string `json:"job_id"`
}

func (o RetryMigrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryMigrationRequest struct{}"
	}

	return strings.Join([]string{"RetryMigrationRequest", string(data)}, " ")
}
