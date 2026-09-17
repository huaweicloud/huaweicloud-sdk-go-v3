package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTopSlowLogRequest Request Object
type ShowInstanceTopSlowLogRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// TOP数量
	Num int32 `json:"num"`

	// 开始时间（Unix timestamp，毫秒）
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp，毫秒）
	EndAt int64 `json:"end_at"`
}

func (o ShowInstanceTopSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTopSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTopSlowLogRequest", string(data)}, " ")
}
