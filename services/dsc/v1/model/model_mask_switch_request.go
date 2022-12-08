package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MaskSwitchRequest struct {

	// 脱敏任务状态
	Status *int32 `json:"status,omitempty"`
}

func (o MaskSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaskSwitchRequest struct{}"
	}

	return strings.Join([]string{"MaskSwitchRequest", string(data)}, " ")
}
