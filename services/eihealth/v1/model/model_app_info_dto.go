package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppInfoDto struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 应用来源项目名称
	AppSrcProjectName *string `json:"app_src_project_name,omitempty"`

	// 应用标签
	AppLabels *[]string `json:"app_labels,omitempty"`

	// 应用简述
	AppSummary *string `json:"app_summary,omitempty"`

	// 应用描述
	AppDescription *string `json:"app_description,omitempty"`

	// 应用镜像
	AppImage *string `json:"app_image,omitempty"`

	// 任务使用到的应用自带的命令信息
	AppCommands *[]string `json:"app_commands,omitempty"`

	// 任务使用到的应用自带的输入参数信息
	AppInputParameters *[]AppInputParameterDto `json:"app_input_parameters,omitempty"`

	// 任务使用到的应用自带的输出参数信息
	AppOutputParameters *[]AppOutputParameterDto `json:"app_output_parameters,omitempty"`

	// 计算节点标签
	AppNodeLabels *[]string `json:"app_node_labels,omitempty"`

	// 图标base64编码
	AppIcon *string `json:"app_icon,omitempty"`
}

func (o AppInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfoDto struct{}"
	}

	return strings.Join([]string{"AppInfoDto", string(data)}, " ")
}
