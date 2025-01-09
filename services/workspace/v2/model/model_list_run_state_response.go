package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRunStateResponse Response Object
type ListRunStateResponse struct {

	// 停止个数。
	StopNum *int32 `json:"stop_num,omitempty"`

	// 运行中个数。
	ActiveNum *int32 `json:"active_num,omitempty"`

	// 故障个数。
	ErrorNum *int32 `json:"error_num,omitempty"`

	// 休眠个数。
	HibernatedNum  *int32 `json:"hibernated_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRunStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunStateResponse struct{}"
	}

	return strings.Join([]string{"ListRunStateResponse", string(data)}, " ")
}
