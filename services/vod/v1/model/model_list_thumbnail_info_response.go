package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThumbnailInfoResponse Response Object
type ListThumbnailInfoResponse struct {

	// 截图结果总数
	Total *int32 `json:"total,omitempty"`

	// 截图结果列表
	Thumbnails     *[]QueryThumbnailInfo `json:"thumbnails,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListThumbnailInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnailInfoResponse struct{}"
	}

	return strings.Join([]string{"ListThumbnailInfoResponse", string(data)}, " ")
}
