package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckImageModerationResponse Response Object
type CheckImageModerationResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	Result         *ImageDetectionResult `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CheckImageModerationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckImageModerationResponse struct{}"
	}

	return strings.Join([]string{"CheckImageModerationResponse", string(data)}, " ")
}
