package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppTemplateResponse Response Object
type ShowAppTemplateResponse struct {

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 模板执行运行时
	Runtime *string `json:"runtime,omitempty"`

	// 模板使用场景
	Category *string `json:"category,omitempty"`

	// 模板参数
	Params *string `json:"params,omitempty"`

	// 模板镜像文件（base64编码）
	Image *string `json:"image,omitempty"`

	// 模板部署次数
	DeployCount *int64 `json:"deploy_count,omitempty"`

	// 模板版本
	Version *int64 `json:"version,omitempty"`

	// 模板指南
	TemplateGuide *string `json:"template_guide,omitempty"`

	// 模板创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 模板更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 模板资源
	Resources      *[]AppTemplateResourceDetail `json:"resources,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowAppTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAppTemplateResponse", string(data)}, " ")
}
