package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingJobInfoRequest Request Object
type ShowSqlLimitingJobInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ShowSqlLimitingJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingJobInfoRequest", string(data)}, " ")
}
