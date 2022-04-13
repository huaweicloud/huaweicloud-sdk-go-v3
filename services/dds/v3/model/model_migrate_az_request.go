package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type MigrateAzRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *MigrateAzRequestBody `json:"body,omitempty"`
}

func (o MigrateAzRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateAzRequest struct{}"
	}

	return strings.Join([]string{"MigrateAzRequest", string(data)}, " ")
}
