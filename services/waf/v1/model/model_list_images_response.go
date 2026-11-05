package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesResponse Response Object
type ListImagesResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 图片列表
	Items          *[]AtlasImageItem `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesResponse struct{}"
	}

	return strings.Join([]string{"ListImagesResponse", string(data)}, " ")
}
