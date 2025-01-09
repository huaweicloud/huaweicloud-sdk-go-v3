package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesSeamless 通用音视频旁路。
type PoliciesSeamless struct {

	// 是否开启通用音视频开关。取值为： false：表示关闭。 true：表示开启。
	SeamlessEnable *bool `json:"seamless_enable,omitempty"`

	Options *PoliciesSeamlessOptions `json:"options,omitempty"`
}

func (o PoliciesSeamless) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesSeamless struct{}"
	}

	return strings.Join([]string{"PoliciesSeamless", string(data)}, " ")
}
