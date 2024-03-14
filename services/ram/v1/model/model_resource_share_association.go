package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceShareAssociation 与资源共享实例关联的资源使用者或共享资源的详细信息。
type ResourceShareAssociation struct {

	// 绑定的实体。这可以是共享资源的URN、账号ID、组织根的URN或OU的URN之一。
	AssociatedEntity string `json:"associated_entity"`

	// 绑定中包含的实体类型。
	AssociationType string `json:"association_type"`

	// 绑定被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 最后一次更新绑定的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 标识资源使用者是否和共享的拥有者属于同一个组织
	External *bool `json:"external,omitempty"`

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	// 资源共享实例的名称。
	ResourceShareName string `json:"resource_share_name"`

	// 绑定的当前状态。
	Status string `json:"status"`

	// 绑定的当前状态的描述。
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o ResourceShareAssociation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceShareAssociation struct{}"
	}

	return strings.Join([]string{"ResourceShareAssociation", string(data)}, " ")
}
