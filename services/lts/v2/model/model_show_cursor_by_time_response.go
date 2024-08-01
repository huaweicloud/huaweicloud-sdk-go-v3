package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCursorByTimeResponse Response Object
type ShowCursorByTimeResponse struct {

	// 游标值
	Cursor         *int64 `json:"cursor,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCursorByTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCursorByTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowCursorByTimeResponse", string(data)}, " ")
}
