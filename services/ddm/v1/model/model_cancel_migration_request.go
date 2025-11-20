package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelMigrationRequest Request Object
type CancelMigrationRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	// 任务流id
	JobId string `json:"job_id"`
}

func (o CancelMigrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelMigrationRequest struct{}"
	}

	return strings.Join([]string{"CancelMigrationRequest", string(data)}, " ")
}
