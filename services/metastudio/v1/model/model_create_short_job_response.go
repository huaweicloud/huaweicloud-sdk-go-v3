package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShortJobResponse Response Object
type CreateShortJobResponse struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 上传训练数据的地址。训练数据需打包成zip文件后，上传至该url。 > 通过该obs地址上传时，需设置content-type为application/zip。
	DataUploadingUrl *string `json:"data_uploading_url,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o CreateShortJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShortJobResponse struct{}"
	}

	return strings.Join([]string{"CreateShortJobResponse", string(data)}, " ")
}
