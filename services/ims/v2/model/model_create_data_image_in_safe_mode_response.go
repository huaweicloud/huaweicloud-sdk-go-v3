package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataImageInSafeModeResponse Response Object
type CreateDataImageInSafeModeResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataImageInSafeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageInSafeModeResponse struct{}"
	}

	return strings.Join([]string{"CreateDataImageInSafeModeResponse", string(data)}, " ")
}
