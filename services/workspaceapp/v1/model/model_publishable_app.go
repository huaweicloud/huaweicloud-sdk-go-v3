package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishableApp 可发布应用。
type PublishableApp struct {

	// 应用ID。
	Id *string `json:"id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 应用版本号。
	Version *string `json:"version,omitempty"`

	// 启动命令行参数。
	CommandParam *string `json:"command_param,omitempty"`

	// 执行路径。
	ExecutePath *string `json:"execute_path,omitempty"`

	// 应用工作目录。
	WorkPath *string `json:"work_path,omitempty"`

	// 应用图标的路径。
	IconPath *string `json:"icon_path,omitempty"`

	// 应用图标的索引。
	IconIndex *int32 `json:"icon_index,omitempty"`

	// 应用描述。
	Description *string `json:"description,omitempty"`

	// 应用组标识Id。
	AppGroupId *string `json:"app_group_id,omitempty"`

	State *AppStateEnum `json:"state,omitempty"`

	// 所在的租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 发布时间。
	PublishAt *sdktime.SdkTime `json:"publish_at,omitempty"`

	// 应用类型： - '1':系统内置应用 - '2':镜像应用 - '3':自定义应用
	SourceType *int32 `json:"source_type,omitempty"`

	// 应用发布者。
	Publisher *string `json:"publisher,omitempty"`

	// 图标url。
	IconUrl *string `json:"icon_url,omitempty"`

	// 是否可发布应用： - true: 可发布。 - false: 不可发布。
	Publishable *bool `json:"publishable,omitempty"`

	// 是否使用沙箱模式运行，取值为： - false: 表示不以沙箱模式运行。 - true: 表示以沙箱模式运行。
	SandboxEnable *bool `json:"sandbox_enable,omitempty"`

	// 镜像ids。
	SourceImageIds *[]string `json:"source_image_ids,omitempty"`
}

func (o PublishableApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishableApp struct{}"
	}

	return strings.Join([]string{"PublishableApp", string(data)}, " ")
}
