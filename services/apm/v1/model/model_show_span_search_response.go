package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSpanSearchResponse struct {

	// 返回的总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// span信息
	SpanInfoList   *[]ClientSpanInfo `json:"span_info_list,omitempty" xml:"span_info_list"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSpanSearchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpanSearchResponse struct{}"
	}

	return strings.Join([]string{"ShowSpanSearchResponse", string(data)}, " ")
}