package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingMiddleJobResponse Response Object
type CreateTrainingMiddleJobResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 上传的训练数据地址,用户需要将训练数据打成zip包后上传到该url,create_type为pakcage时设置。 > * 通过该obs地址上传时需设置content-type为application/zip
	TrainingDataUploadingUrl *string `json:"training_data_uploading_url,omitempty"`

	SegmentUploadingUrl *CreateTrainingJobRspSegmentUploadingUrl `json:"segment_uploading_url,omitempty"`

	// 授权书的上传地址。
	AuthorizationLetterUploadingUrl *string `json:"authorization_letter_uploading_url,omitempty"`
	HttpStatusCode                  int     `json:"-"`
}

func (o CreateTrainingMiddleJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingMiddleJobResponse struct{}"
	}

	return strings.Join([]string{"CreateTrainingMiddleJobResponse", string(data)}, " ")
}
