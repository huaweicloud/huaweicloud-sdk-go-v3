package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThumbnailDetailsResponse Response Object
type ListThumbnailDetailsResponse struct {

	// 截图图片总数
	Total *int32 `json:"total,omitempty"`

	// 截图结果列表
	ThumbnailResult *[]ThumbnailRsp `json:"thumbnail_result,omitempty"`
	HttpStatusCode  int             `json:"-"`
}

func (o ListThumbnailDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnailDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListThumbnailDetailsResponse", string(data)}, " ")
}
