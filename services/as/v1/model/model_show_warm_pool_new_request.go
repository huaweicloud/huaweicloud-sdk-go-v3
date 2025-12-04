package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWarmPoolNewRequest Request Object
type ShowWarmPoolNewRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowWarmPoolNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWarmPoolNewRequest struct{}"
	}

	return strings.Join([]string{"ShowWarmPoolNewRequest", string(data)}, " ")
}
