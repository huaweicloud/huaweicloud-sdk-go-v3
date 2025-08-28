package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWarmPoolRequest Request Object
type ShowWarmPoolRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowWarmPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWarmPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowWarmPoolRequest", string(data)}, " ")
}
