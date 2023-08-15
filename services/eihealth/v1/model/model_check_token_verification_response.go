package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTokenVerificationResponse Response Object
type CheckTokenVerificationResponse struct {

	// 过期时间
	ExpiresTime *sdktime.SdkTime `json:"expires_time,omitempty"`

	Project *ProjectDto `json:"project,omitempty"`

	// 角色
	Roles *[]RoleDto `json:"roles,omitempty"`

	User           *UserDto `json:"user,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CheckTokenVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTokenVerificationResponse struct{}"
	}

	return strings.Join([]string{"CheckTokenVerificationResponse", string(data)}, " ")
}
