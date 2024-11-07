package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollBackDatabaseVersionResponse Response Object
type RollBackDatabaseVersionResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RollBackDatabaseVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollBackDatabaseVersionResponse struct{}"
	}

	return strings.Join([]string{"RollBackDatabaseVersionResponse", string(data)}, " ")
}
