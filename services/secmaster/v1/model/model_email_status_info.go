package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EmailStatusInfo 邮箱状态信息
type EmailStatusInfo struct {

	// 邮箱地址
	Address string `json:"address"`

	// 邮箱状态，FINISH表示已通过校验；其他状态需要校验
	ConfirmStatus string `json:"confirm_status"`
}

func (o EmailStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailStatusInfo struct{}"
	}

	return strings.Join([]string{"EmailStatusInfo", string(data)}, " ")
}
