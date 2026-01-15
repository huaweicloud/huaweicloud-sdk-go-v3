package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesUserProfile UPM配置。
type PoliciesUserProfile struct {

	// 是否开启UPM配置。取值为： false：表示关闭。 true：表示开启。
	UserProfileEnable *bool `json:"user_profile_enable,omitempty"`

	Options *PoliciesUserProfileOptions `json:"options,omitempty"`
}

func (o PoliciesUserProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesUserProfile struct{}"
	}

	return strings.Join([]string{"PoliciesUserProfile", string(data)}, " ")
}
