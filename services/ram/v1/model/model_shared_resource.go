package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SharedResource 描述RAM中的资源共享关联的资源。
type SharedResource struct {

	// 资源的统一资源名称。
	ResourceUrn *string `json:"resource_urn,omitempty"`

	// 资源与资源共享实例关联的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 最后一次更新资源共享实例的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 资源的类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源绑定的资源共享实例的ID。
	ResourceShareId *string `json:"resource_share_id,omitempty"`

	// 资源绑定的当前状态。
	Status *string `json:"status,omitempty"`
}

func (o SharedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SharedResource struct{}"
	}

	return strings.Join([]string{"SharedResource", string(data)}, " ")
}
