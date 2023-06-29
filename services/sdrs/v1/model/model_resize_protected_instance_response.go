package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeProtectedInstanceResponse Response Object
type ResizeProtectedInstanceResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeProtectedInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceResponse", string(data)}, " ")
}
