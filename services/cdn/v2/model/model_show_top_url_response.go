package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTopUrlResponse struct {

	// 服务区域
	ServiceArea *string `json:"service_area,omitempty" xml:"service_area"`

	// 详情数据对象。
	TopUrlSummary  *[]TopUrlSummary `json:"top_url_summary,omitempty" xml:"top_url_summary"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTopUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowTopUrlResponse", string(data)}, " ")
}
