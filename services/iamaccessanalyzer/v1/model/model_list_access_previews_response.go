package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPreviewsResponse Response Object
type ListAccessPreviewsResponse struct {

	// 访问预览列表。
	AccessPreviews *[]AccessPreviewSummary `json:"access_previews,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAccessPreviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPreviewsResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPreviewsResponse", string(data)}, " ")
}
