package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityProfilesResponse Response Object
type ListSecurityProfilesResponse struct {

	// 安全态势感知配置信息列表。
	SecurityProfiles *[]SecurityProfileDto `json:"security_profiles,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListSecurityProfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityProfilesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityProfilesResponse", string(data)}, " ")
}
