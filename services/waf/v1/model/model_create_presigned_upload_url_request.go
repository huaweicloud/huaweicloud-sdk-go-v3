package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePresignedUploadUrlRequest Request Object
type CreatePresignedUploadUrlRequest struct {

	// 图库ID
	AtlasId string `json:"atlas_id"`

	Body *PresignUploadUrlBody `json:"body,omitempty"`
}

func (o CreatePresignedUploadUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePresignedUploadUrlRequest struct{}"
	}

	return strings.Join([]string{"CreatePresignedUploadUrlRequest", string(data)}, " ")
}
