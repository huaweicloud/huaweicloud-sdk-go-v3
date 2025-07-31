package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RaspProtectStatus 动态网页防篡改防护状态   - opened : 防护中   - closed : 未防护
type RaspProtectStatus struct {
}

func (o RaspProtectStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RaspProtectStatus struct{}"
	}

	return strings.Join([]string{"RaspProtectStatus", string(data)}, " ")
}
