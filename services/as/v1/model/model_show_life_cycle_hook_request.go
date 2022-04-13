package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLifeCycleHookRequest struct {
	// 伸缩组标识。

	ScalingGroupId string `json:"scaling_group_id"`
	// 生命周期挂钩标识。

	LifecycleHookName string `json:"lifecycle_hook_name"`
}

func (o ShowLifeCycleHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLifeCycleHookRequest struct{}"
	}

	return strings.Join([]string{"ShowLifeCycleHookRequest", string(data)}, " ")
}
