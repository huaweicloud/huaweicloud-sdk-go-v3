package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseWarmPoolRequest Request Object
type CloseWarmPoolRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o CloseWarmPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseWarmPoolRequest struct{}"
	}

	return strings.Join([]string{"CloseWarmPoolRequest", string(data)}, " ")
}
