package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanMigrationRequest Request Object
type CleanMigrationRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	// 任务流id
	JobId string `json:"job_id"`
}

func (o CleanMigrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanMigrationRequest struct{}"
	}

	return strings.Join([]string{"CleanMigrationRequest", string(data)}, " ")
}
