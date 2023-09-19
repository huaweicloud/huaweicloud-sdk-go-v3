package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePhotoDigitalHumanVideoResponse Response Object
type CreatePhotoDigitalHumanVideoResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePhotoDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePhotoDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"CreatePhotoDigitalHumanVideoResponse", string(data)}, " ")
}
