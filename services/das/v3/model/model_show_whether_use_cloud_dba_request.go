package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWhetherUseCloudDbaRequest Request Object
type ShowWhetherUseCloudDbaRequest struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ShowWhetherUseCloudDbaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhetherUseCloudDbaRequest struct{}"
	}

	return strings.Join([]string{"ShowWhetherUseCloudDbaRequest", string(data)}, " ")
}
