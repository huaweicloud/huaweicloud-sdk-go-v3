package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrrTemplateResponse Response Object
type ListPrrTemplateResponse struct {

	// PRR模板id
	Id *string `json:"id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 应用分类
	ApplicationType *string `json:"application_type,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 创建人id
	Creator *string `json:"creator,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最后更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 是否已关联检查项
	IsRelatedReview *bool `json:"is_related_review,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListPrrTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrrTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListPrrTemplateResponse", string(data)}, " ")
}
