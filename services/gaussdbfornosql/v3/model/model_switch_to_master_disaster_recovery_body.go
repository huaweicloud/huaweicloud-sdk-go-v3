package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchToMasterDisasterRecoveryBody struct {

	// 是否强制备实例升主。 若为True，则强制备实例升主，用于在主实例异常的状态下，快速恢复服务的场景：允许备实例强制升为特殊主实例，独立提供读写服务。 默认为False，用于正常状态下备实例平缓升主。
	Force *bool `json:"force,omitempty"`
}

func (o SwitchToMasterDisasterRecoveryBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToMasterDisasterRecoveryBody struct{}"
	}

	return strings.Join([]string{"SwitchToMasterDisasterRecoveryBody", string(data)}, " ")
}
