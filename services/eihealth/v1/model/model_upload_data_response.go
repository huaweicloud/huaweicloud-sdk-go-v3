package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadDataResponse struct {

	// 分段上传任务id
	MultipartId    *string `json:"multipart_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDataResponse struct{}"
	}

	return strings.Join([]string{"UploadDataResponse", string(data)}, " ")
}
