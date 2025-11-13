package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserReviewAttachmentUploadingAddressRequest Request Object
type ShowUserReviewAttachmentUploadingAddressRequest struct {

	// 任务id
	JobId string `json:"job_id"`

	// 起始序号
	StartNumber *int32 `json:"start_number,omitempty"`

	// 结束序号
	EndNumber *int32 `json:"end_number,omitempty"`
}

func (o ShowUserReviewAttachmentUploadingAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserReviewAttachmentUploadingAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowUserReviewAttachmentUploadingAddressRequest", string(data)}, " ")
}
