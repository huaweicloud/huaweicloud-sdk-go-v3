package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckMigrateLogicDbRequest Request Object
type CheckMigrateLogicDbRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	Body *MigrateLogicDbOpenReq `json:"body,omitempty"`
}

func (o CheckMigrateLogicDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMigrateLogicDbRequest struct{}"
	}

	return strings.Join([]string{"CheckMigrateLogicDbRequest", string(data)}, " ")
}
