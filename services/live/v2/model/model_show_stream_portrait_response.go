package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStreamPortraitResponse struct {

	// 播放画像信息列表。
	StreamPortraits *[]StreamPortrait `json:"stream_portraits,omitempty" xml:"stream_portraits"`

	XRequestId     *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStreamPortraitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamPortraitResponse struct{}"
	}

	return strings.Join([]string{"ShowStreamPortraitResponse", string(data)}, " ")
}
