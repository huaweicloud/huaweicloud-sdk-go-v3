package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserProfileManagementOptions 用户漫游配置。
type UserProfileManagementOptions struct {
	Profiles *UserProfileManagementProfiles `json:"profiles,omitempty"`
}

func (o UserProfileManagementOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserProfileManagementOptions struct{}"
	}

	return strings.Join([]string{"UserProfileManagementOptions", string(data)}, " ")
}
