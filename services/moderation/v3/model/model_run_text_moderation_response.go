package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunTextModerationResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存
	RequestId *string `json:"request_id,omitempty"`

	Result         *TextDetectionResult `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RunTextModerationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextModerationResponse struct{}"
	}

	return strings.Join([]string{"RunTextModerationResponse", string(data)}, " ")
}
