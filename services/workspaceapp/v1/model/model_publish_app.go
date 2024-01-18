package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApp 发布应用的结构体。
type PublishApp struct {

	// 应用名称,名称需满足如下规则: 1. 名称允许可见字符或空格，不可为全空格。 2. 不允许包含如下字符:^;|~`{}[]<>。 3. 长度1~64个字符。
	Name string `json:"name"`

	// 应用版本号。
	Version *string `json:"version,omitempty"`

	// 启动命令行参数。
	CommandParam *string `json:"command_param,omitempty"`

	// 图标地址，该字段当前未使用。 > - 图片的默认大小当前限制为8KB，即1024 * 8字节。 > - 如果数据格式为data;image/png;base64,iVBORw0KGgoAAAANS时实际大小约为字段约为：size * 4/3 + 4bytes。
	IconUri *string `json:"icon_uri,omitempty"`

	// 执行路径。
	ExecutePath string `json:"execute_path"`

	// 应用工作目录。
	WorkPath *string `json:"work_path,omitempty"`

	// 应用图标的路径。
	IconPath *string `json:"icon_path,omitempty"`

	// 应用图标的索引。
	IconIndex *int32 `json:"icon_index,omitempty"`

	// 应用描述。
	Description *string `json:"description,omitempty"`

	// 应用类型： - '1':系统保留不可用。 - '2':镜像应用。 - '3':自定义应用。
	SourceType int32 `json:"source_type"`

	// 应用发布者。
	Publisher *string `json:"publisher,omitempty"`

	// 镜像ids,最多20个。
	SourceImageIds *[]string `json:"source_image_ids,omitempty"`

	// 是否使用沙箱模式运行，取值为： - false: 表示不以沙箱模式运行。 - true: 表示以沙箱模式运行。
	SandboxEnable *bool `json:"sandbox_enable,omitempty"`
}

func (o PublishApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApp struct{}"
	}

	return strings.Join([]string{"PublishApp", string(data)}, " ")
}
