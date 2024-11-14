package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBeautyPreviewJobResponse Response Object
type CreateBeautyPreviewJobResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 美白前图片上传url。
	PreBeautyImageUploadUrl *string `json:"pre_beauty_image_upload_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBeautyPreviewJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBeautyPreviewJobResponse struct{}"
	}

	return strings.Join([]string{"CreateBeautyPreviewJobResponse", string(data)}, " ")
}
