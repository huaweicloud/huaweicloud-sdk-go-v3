package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesRequest Request Object
type ListImagesRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`

	// 分页偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数，1~1000，默认10
	Limit *int32 `json:"limit,omitempty"`

	// 图片名称（模糊匹配）
	Name *string `json:"name,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}
