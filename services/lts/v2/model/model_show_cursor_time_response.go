package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCursorTimeResponse Response Object
type ShowCursorTimeResponse struct {

	// 游标时间值
	CursorTime     *int64 `json:"cursor_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCursorTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCursorTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowCursorTimeResponse", string(data)}, " ")
}
