package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTokenVerificationResponse struct {

	// 过期时间
	ExpiresTime *sdktime.SdkTime `json:"expires_time,omitempty"`

	Project *ProjectDto `json:"project,omitempty"`

	// 角色
	Roles *[]RoleDto `json:"roles,omitempty"`

	User           *UserDto `json:"user,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowTokenVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenVerificationResponse struct{}"
	}

	return strings.Join([]string{"ShowTokenVerificationResponse", string(data)}, " ")
}
