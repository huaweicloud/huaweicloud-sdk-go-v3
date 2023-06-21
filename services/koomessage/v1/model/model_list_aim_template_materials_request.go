package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAimTemplateMaterialsRequest struct {

	// 资源类型。 - image：表示图片 - video：表示视频
	ResourceType string `json:"resource_type"`

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 素材ID。
	MaterialId *string `json:"material_id,omitempty"`

	// 资源ID。
	AimResourceId *string `json:"aim_resource_id,omitempty"`

	// 翻页页数，从1开始。
	Offset int32 `json:"offset"`

	// 每页展示的条数。
	Limit int32 `json:"limit"`
}

func (o ListAimTemplateMaterialsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateMaterialsRequest struct{}"
	}

	return strings.Join([]string{"ListAimTemplateMaterialsRequest", string(data)}, " ")
}
