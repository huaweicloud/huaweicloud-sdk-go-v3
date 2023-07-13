package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RateControl struct {

	// 是否分批发布，默认值是false。
	HaveRateControl *bool `json:"have_rate_control,omitempty"`

	// 每批间隔。
	TimeDelay *int32 `json:"time_delay,omitempty"`

	// 每批支持的最大实例数。
	Max *int32 `json:"max,omitempty"`
}

func (o RateControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RateControl struct{}"
	}

	return strings.Join([]string{"RateControl", string(data)}, " ")
}
