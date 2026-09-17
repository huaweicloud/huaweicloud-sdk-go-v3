package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopSlowLogRequest Request Object
type ShowTopSlowLogRequest struct {

	// TOP数量
	Num int32 `json:"num"`

	// 开始时间（Unix timestamp，毫秒）
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp，毫秒）
	EndAt int64 `json:"end_at"`
}

func (o ShowTopSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ShowTopSlowLogRequest", string(data)}, " ")
}
