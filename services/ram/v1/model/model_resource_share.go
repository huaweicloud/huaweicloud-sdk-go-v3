package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 描述RAM中的资源共享。
type ResourceShare struct {

	// 资源共享实例的ID。
	Id string `json:"id"`

	// 资源共享实例的名称。
	Name string `json:"name"`

	// 资源共享实例的描述。
	Description string `json:"description"`

	// 资源共享实例的所有者ID。
	OwningAccountId string `json:"owning_account_id"`

	// 资源共享实例的状态。
	Status string `json:"status"`

	// 资源共享实例的标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 创建资源共享实例的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 最后一次更新资源共享实例的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o ResourceShare) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceShare struct{}"
	}

	return strings.Join([]string{"ResourceShare", string(data)}, " ")
}
