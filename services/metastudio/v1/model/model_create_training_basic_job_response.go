package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingBasicJobResponse Response Object
type CreateTrainingBasicJobResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 上传训练数据的地址。训练数据需打包成zip文件后，上传至该url。  create_type取值为package时设置。 > 通过该obs地址上传时，需设置content-type为application/zip。
	TrainingDataUploadingUrl *string `json:"training_data_uploading_url,omitempty"`

	SegmentUploadingUrl *CreateTrainingJobRspSegmentUploadingUrl `json:"segment_uploading_url,omitempty"`

	// 授权书的上传地址。
	AuthorizationLetterUploadingUrl *string `json:"authorization_letter_uploading_url,omitempty"`
	HttpStatusCode                  int     `json:"-"`
}

func (o CreateTrainingBasicJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingBasicJobResponse struct{}"
	}

	return strings.Join([]string{"CreateTrainingBasicJobResponse", string(data)}, " ")
}
