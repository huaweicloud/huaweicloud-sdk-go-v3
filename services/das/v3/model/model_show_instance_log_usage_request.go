package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceLogUsageRequest Request Object
type ShowInstanceLogUsageRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceLogUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceLogUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceLogUsageRequest", string(data)}, " ")
}
