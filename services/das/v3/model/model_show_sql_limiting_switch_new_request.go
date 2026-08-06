package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingSwitchNewRequest Request Object
type ShowSqlLimitingSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ShowSqlLimitingSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingSwitchNewRequest", string(data)}, " ")
}
