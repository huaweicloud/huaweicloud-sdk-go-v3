package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptanceSchema the request body of acceptance
type AcceptanceSchema struct {

	// request_id
	RequestId string `json:"request_id"`

	// task_id
	TaskId int32 `json:"task_id"`
}

func (o AcceptanceSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptanceSchema struct{}"
	}

	return strings.Join([]string{"AcceptanceSchema", string(data)}, " ")
}
