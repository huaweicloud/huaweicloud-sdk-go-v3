package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePhotoDetectionResponse Response Object
type CreatePhotoDetectionResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePhotoDetectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePhotoDetectionResponse struct{}"
	}

	return strings.Join([]string{"CreatePhotoDetectionResponse", string(data)}, " ")
}
