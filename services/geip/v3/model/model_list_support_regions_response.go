package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportRegionsResponse Response Object
type ListSupportRegionsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 支持的Region对象
	SupportRegions *[]ListSupportRegions `json:"support_regions,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSupportRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportRegionsResponse", string(data)}, " ")
}
