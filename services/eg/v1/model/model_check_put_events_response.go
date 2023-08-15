package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPutEventsResponse Response Object
type CheckPutEventsResponse struct {

	// 预校验发布事件失败的个数
	FailedCount *int32 `json:"failed_count,omitempty"`

	Sources *[]CheckPutEventsResult `json:"sources,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckPutEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPutEventsResponse struct{}"
	}

	return strings.Join([]string{"CheckPutEventsResponse", string(data)}, " ")
}
