package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserProfileManagementProfiles 用户漫游profiles配置。
type UserProfileManagementProfiles struct {

	// 配置VHD会话回写状态： 0: 已禁用 1: 已启用 2: 未配置
	VhdWriteBack *string `json:"vhd_write_back,omitempty"`
}

func (o UserProfileManagementProfiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserProfileManagementProfiles struct{}"
	}

	return strings.Join([]string{"UserProfileManagementProfiles", string(data)}, " ")
}
