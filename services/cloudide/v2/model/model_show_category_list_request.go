package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCategoryListRequest Request Object
type ShowCategoryListRequest struct {

	// 页码
	PageNum *int64 `json:"page_num,omitempty"`

	// 分页大小
	PageSize *int64 `json:"page_size,omitempty"`

	// 场景名称
	SceneName *string `json:"scene_name,omitempty"`

	// 支持的ide
	SupportIde *int32 `json:"support_ide,omitempty"`

	// 数据来源
	Type *int32 `json:"type,omitempty"`
}

func (o ShowCategoryListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCategoryListRequest struct{}"
	}

	return strings.Join([]string{"ShowCategoryListRequest", string(data)}, " ")
}
