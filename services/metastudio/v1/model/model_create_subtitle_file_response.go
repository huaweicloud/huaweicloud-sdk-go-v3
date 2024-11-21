package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubtitleFileResponse Response Object
type CreateSubtitleFileResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubtitleFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubtitleFileResponse struct{}"
	}

	return strings.Join([]string{"CreateSubtitleFileResponse", string(data)}, " ")
}
