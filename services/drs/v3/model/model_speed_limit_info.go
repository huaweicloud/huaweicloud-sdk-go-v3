package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 限制迁移速度请求体
type SpeedLimitInfo struct {

	// 开始限速时间, 此时间为UTC时间，开始时间为整时，若有分钟，则会忽略，格式为hh:mm，小时数为两位，例如：01:00。
	Begin string `json:"begin"`

	// 结束时间,此时间为UTC时间,输入必须为59分结尾，格式为hh:mm，小时数为两位，例如：05:59。
	End string `json:"end"`

	// 限速，取值范围为1~9999 ,单位为MB/s
	Speed string `json:"speed"`

	// 是否为UTC时间
	IsUtc *bool `json:"is_utc,omitempty"`
}

func (o SpeedLimitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeedLimitInfo struct{}"
	}

	return strings.Join([]string{"SpeedLimitInfo", string(data)}, " ")
}
