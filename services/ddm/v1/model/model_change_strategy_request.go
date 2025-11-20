package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeStrategyRequest Request Object
type ChangeStrategyRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	// 任务流id
	JobId string `json:"job_id"`

	Body *MigrateRouteSwitchReqVo `json:"body,omitempty"`
}

func (o ChangeStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeStrategyRequest struct{}"
	}

	return strings.Join([]string{"ChangeStrategyRequest", string(data)}, " ")
}
