package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalEipTagsResponse Response Object
type ShowGlobalEipTagsResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 单个资源的租户标签列表。
	Tags *[]CreateTag `json:"tags,omitempty"`

	// 单个资源的系统标签列表。
	SysTags *[]SysTag `json:"sys_tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGlobalEipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalEipTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalEipTagsResponse", string(data)}, " ")
}
