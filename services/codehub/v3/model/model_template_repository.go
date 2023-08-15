package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateRepository struct {

	// 模板唯一标识
	Id *int32 `json:"id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板关联仓库名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板标签
	Tags *[]string `json:"tags,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 模板简介
	BriefIntroduction *string `json:"brief_introduction,omitempty"`

	// 是否自动创建流水线
	AutoPendingPipelines *int32 `json:"auto_pending_pipelines,omitempty"`

	// 模板语言分类
	Language *string `json:"language,omitempty"`

	// 模板创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 模板引用次数
	UsedTimes *int32 `json:"used_times,omitempty"`

	// 模板被点赞次数
	LikedTimes *int32 `json:"liked_times,omitempty"`

	// 模板创建人
	CreatorName *string `json:"creator_name,omitempty"`

	// 模板https链接
	HttpsUrl *string `json:"https_url,omitempty"`
}

func (o TemplateRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRepository struct{}"
	}

	return strings.Join([]string{"TemplateRepository", string(data)}, " ")
}
