package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotSetChargeModeInstanceRequest Request Object
type ListNotSetChargeModeInstanceRequest struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ListNotSetChargeModeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotSetChargeModeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListNotSetChargeModeInstanceRequest", string(data)}, " ")
}
