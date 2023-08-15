package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHighRiskCommandsRequest Request Object
type ShowHighRiskCommandsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowHighRiskCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHighRiskCommandsRequest struct{}"
	}

	return strings.Join([]string{"ShowHighRiskCommandsRequest", string(data)}, " ")
}
