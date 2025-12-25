package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCodeSegmentResponse Response Object
type DeleteCodeSegmentResponse struct {

	// UUID
	CodeSegmentId *string `json:"code_segment_id,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteCodeSegmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCodeSegmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteCodeSegmentResponse", string(data)}, " ")
}
