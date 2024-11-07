package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDatabaseVersionRequest Request Object
type ChangeDatabaseVersionRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *DatabaseVersionRequest `json:"body,omitempty"`
}

func (o ChangeDatabaseVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDatabaseVersionRequest struct{}"
	}

	return strings.Join([]string{"ChangeDatabaseVersionRequest", string(data)}, " ")
}
