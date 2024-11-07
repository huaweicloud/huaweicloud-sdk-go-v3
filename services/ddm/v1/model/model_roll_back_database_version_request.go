package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollBackDatabaseVersionRequest Request Object
type RollBackDatabaseVersionRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`
}

func (o RollBackDatabaseVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollBackDatabaseVersionRequest struct{}"
	}

	return strings.Join([]string{"RollBackDatabaseVersionRequest", string(data)}, " ")
}
