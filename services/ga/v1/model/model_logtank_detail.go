package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogtankDetail 云日志配置详情。
type LogtankDetail struct {

	// 云日志ID。
	Id *string `json:"id,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	ResourceType *LogtankResourceType `json:"resource_type,omitempty"`

	// 开启云日志的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 云日志组ID。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 云日志流ID。
	LogStreamId *string `json:"log_stream_id,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o LogtankDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogtankDetail struct{}"
	}

	return strings.Join([]string{"LogtankDetail", string(data)}, " ")
}
