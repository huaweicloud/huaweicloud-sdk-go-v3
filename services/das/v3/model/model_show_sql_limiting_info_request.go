package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingInfoRequest Request Object
type ShowSqlLimitingInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o ShowSqlLimitingInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingInfoRequest", string(data)}, " ")
}
