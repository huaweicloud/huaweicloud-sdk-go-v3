package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupportKeyStringRequest Request Object
type ShowSupportKeyStringRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`
}

func (o ShowSupportKeyStringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportKeyStringRequest struct{}"
	}

	return strings.Join([]string{"ShowSupportKeyStringRequest", string(data)}, " ")
}
