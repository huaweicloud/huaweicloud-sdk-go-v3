package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceInfoRequest Request Object
type ShowInstanceInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o ShowInstanceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceInfoRequest", string(data)}, " ")
}
