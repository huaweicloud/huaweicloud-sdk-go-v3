package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClickHouseInstanceResponse Response Object
type ResizeClickHouseInstanceResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeClickHouseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClickHouseInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeClickHouseInstanceResponse", string(data)}, " ")
}
