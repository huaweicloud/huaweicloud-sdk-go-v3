package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableRdsForMigrateRequest Request Object
type ListAvailableRdsForMigrateRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`
}

func (o ListAvailableRdsForMigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsForMigrateRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsForMigrateRequest", string(data)}, " ")
}
