package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SharedPrincipal 描述资源共享管理的角色
type SharedPrincipal struct {

	// 资源使用者绑定的资源共享实例的ID。
	ResourceShareId *string `json:"resource_share_id,omitempty"`

	// 资源使用者的账号ID或URN。
	Id *string `json:"id,omitempty"`

	// 资源使用者与资源共享实例关联的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 最后一次更新资源共享实例的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o SharedPrincipal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SharedPrincipal struct{}"
	}

	return strings.Join([]string{"SharedPrincipal", string(data)}, " ")
}
