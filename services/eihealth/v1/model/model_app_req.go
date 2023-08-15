package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppReq 应用请求体
type AppReq struct {

	// 应用名称 目标应用名称 取值范围：长度为[1,56]，以大小写字母开头，允许出现中划线(-)、下划线(_)、小写字母和数字，且必须以大小写字母或数字结尾。
	Name string `json:"name"`

	// 应用版本 取值范围：[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。更新应用时，应用版本不支持修改。
	Version string `json:"version"`

	// 应用简述 取值范围[0,128]
	Summary *string `json:"summary,omitempty"`

	// 应用描述 取值范围[0,65535]，后续支持markdown文本
	Description *string `json:"description,omitempty"`

	// 应用标签 取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者结尾。
	Labels *[]string `json:"labels,omitempty"`

	// docker镜像地址 取值范围[5-255]，不能包含中文字符
	Image string `json:"image"`

	// docker启动时执行命令 取值范围[1-300]，单个命令长度取值范围[0-256]，不能包含中文字符
	Commands []string `json:"commands"`

	Resources *ResourceDto `json:"resources,omitempty"`

	// 应用的输入参数列表
	Inputs *[]AppInputParameterDto `json:"inputs,omitempty"`

	// 应用的输出参数列表
	Outputs *[]AppOutputParameterDto `json:"outputs,omitempty"`

	// 节点标签 取值范围[0,1]，单个标签最大长度63字符
	NodeLabels *[]string `json:"node_labels,omitempty"`

	// 图标base64编码
	Icon *string `json:"icon,omitempty"`
}

func (o AppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppReq struct{}"
	}

	return strings.Join([]string{"AppReq", string(data)}, " ")
}
