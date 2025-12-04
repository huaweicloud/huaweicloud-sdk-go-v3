package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseWarmPoolNewRequest Request Object
type CloseWarmPoolNewRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o CloseWarmPoolNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseWarmPoolNewRequest struct{}"
	}

	return strings.Join([]string{"CloseWarmPoolNewRequest", string(data)}, " ")
}
