package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesUserProfileManagement 用户配置策略。
type PoliciesUserProfileManagement struct {

	// 用户配置状态： 0: 已禁用 1: 已启用 2: 未配置
	UpmStatus *int32 `json:"upm_status,omitempty"`

	Options *UserProfileManagementOptions `json:"options,omitempty"`
}

func (o PoliciesUserProfileManagement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesUserProfileManagement struct{}"
	}

	return strings.Join([]string{"PoliciesUserProfileManagement", string(data)}, " ")
}
