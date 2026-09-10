package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskInfoRequest Request Object
type ShowRiskInfoRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ShowRiskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowRiskInfoRequest", string(data)}, " ")
}
