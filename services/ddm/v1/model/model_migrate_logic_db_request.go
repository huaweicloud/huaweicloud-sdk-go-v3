package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateLogicDbRequest Request Object
type MigrateLogicDbRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	Body *MigrateLogicDbOpenReq `json:"body,omitempty"`
}

func (o MigrateLogicDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateLogicDbRequest struct{}"
	}

	return strings.Join([]string{"MigrateLogicDbRequest", string(data)}, " ")
}
