package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoLogoutOptions struct {

	// 会话断连保留时长（分钟）
	SbcLogoutWaitingTime *int32 `json:"sbc_logout_waiting_time,omitempty"`
}

func (o AutoLogoutOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLogoutOptions struct{}"
	}

	return strings.Join([]string{"AutoLogoutOptions", string(data)}, " ")
}
