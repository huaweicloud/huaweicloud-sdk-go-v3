package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LockInfo struct {

	// 锁id
	Id *string `json:"id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 场景类型
	Scene *string `json:"scene,omitempty"`

	// 源类型
	SourceType *string `json:"source_type,omitempty"`

	// 源id
	SourceId *string `json:"source_id,omitempty"`

	// check地址
	CheckUrl *string `json:"check_url,omitempty"`

	// 动作类型
	Action *string `json:"action,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o LockInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockInfo struct{}"
	}

	return strings.Join([]string{"LockInfo", string(data)}, " ")
}
