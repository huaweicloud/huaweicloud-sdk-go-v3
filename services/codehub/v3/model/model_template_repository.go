package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateRepository struct {

	// 模板唯一标识
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 模板名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 模板关联仓库名称
	TemplateName *string `json:"template_name,omitempty" xml:"template_name"`

	// 模板标签
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 模板描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 模板简介
	BriefIntroduction *string `json:"brief_introduction,omitempty" xml:"brief_introduction"`

	// 是否自动创建流水线
	AutoPendingPipelines *int32 `json:"auto_pending_pipelines,omitempty" xml:"auto_pending_pipelines"`

	// 模板语言分类
	Language *string `json:"language,omitempty" xml:"language"`

	// 模板创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 模板引用次数
	UsedTimes *int32 `json:"used_times,omitempty" xml:"used_times"`

	// 模板被点赞次数
	LikedTimes *int32 `json:"liked_times,omitempty" xml:"liked_times"`

	// 模板创建人
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 模板https链接
	HttpsUrl *string `json:"https_url,omitempty" xml:"https_url"`
}

func (o TemplateRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRepository struct{}"
	}

	return strings.Join([]string{"TemplateRepository", string(data)}, " ")
}
