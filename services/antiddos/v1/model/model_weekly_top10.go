package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeeklyTop10
type WeeklyTop10 struct {

	// 弹性IP地址
	FloatingIpAddress string `json:"floating_ip_address"`

	// DDoS拦截次数，包括清洗和黑洞
	Times int32 `json:"times"`
}

func (o WeeklyTop10) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeeklyTop10 struct{}"
	}

	return strings.Join([]string{"WeeklyTop10", string(data)}, " ")
}
