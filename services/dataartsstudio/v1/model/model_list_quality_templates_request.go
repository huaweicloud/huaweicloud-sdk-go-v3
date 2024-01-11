package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityTemplatesRequest Request Object
type ListQualityTemplatesRequest struct {

	// category id
	CategoryId *int64 `json:"category_id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 是否只查询系统模板
	SystemTemplate *bool `json:"system_template,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 分页时每页的条数,最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ListQualityTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListQualityTemplatesRequest", string(data)}, " ")
}
