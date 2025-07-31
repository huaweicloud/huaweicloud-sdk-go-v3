package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageInSafeModeResponse Response Object
type CreateImageInSafeModeResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateImageInSafeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageInSafeModeResponse struct{}"
	}

	return strings.Join([]string{"CreateImageInSafeModeResponse", string(data)}, " ")
}
