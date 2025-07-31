package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WtpStatus 防护状态   - closed : 未防护   - opened : 防护中   - open_failed : 防护失败   - opening : 正在开启   - partial_protection : 部分防护
type WtpStatus struct {
}

func (o WtpStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpStatus struct{}"
	}

	return strings.Join([]string{"WtpStatus", string(data)}, " ")
}
