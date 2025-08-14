package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesUserProfile 配置文件漫游技术配置。
type PoliciesUserProfile struct {

	// 配置文件漫游技术配置开关： false: 表示关闭 true: 表示开启
	UserProfileEnable *bool `json:"user_profile_enable,omitempty"`

	Options *UpmOptions `json:"options,omitempty"`
}

func (o PoliciesUserProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesUserProfile struct{}"
	}

	return strings.Join([]string{"PoliciesUserProfile", string(data)}, " ")
}
