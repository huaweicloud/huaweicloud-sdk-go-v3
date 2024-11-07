package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDatabaseVersionResponse Response Object
type ChangeDatabaseVersionResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDatabaseVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDatabaseVersionResponse struct{}"
	}

	return strings.Join([]string{"ChangeDatabaseVersionResponse", string(data)}, " ")
}
