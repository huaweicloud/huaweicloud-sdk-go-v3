package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadDbObjectTemplateResponse Response Object
type UploadDbObjectTemplateResponse struct {

	// 是否上传完成。
	SendSuccess *bool `json:"send_success,omitempty"`

	// 文件解析状态。 取值：success，failed
	ProcessStatus *string `json:"process_status,omitempty"`

	// 解析成功的行数。
	ParsedSuccessNumber *string `json:"parsed_success_number,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UploadDbObjectTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDbObjectTemplateResponse struct{}"
	}

	return strings.Join([]string{"UploadDbObjectTemplateResponse", string(data)}, " ")
}
