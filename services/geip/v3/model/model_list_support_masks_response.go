package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportMasksResponse Response Object
type ListSupportMasksResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 支持全域弹性公网IP段的掩码范围列表
	SupportMasks *[]ListSupportMasks `json:"support_masks,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSupportMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportMasksResponse struct{}"
	}

	return strings.Join([]string{"ListSupportMasksResponse", string(data)}, " ")
}
