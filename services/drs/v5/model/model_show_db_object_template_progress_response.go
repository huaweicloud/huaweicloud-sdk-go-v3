package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDbObjectTemplateProgressResponse Response Object
type ShowDbObjectTemplateProgressResponse struct {

	// 是否上传完成。
	SendSuccess *bool `json:"send_success,omitempty"`

	// 文件解析状态。
	ProcessStatus *string `json:"process_status,omitempty"`

	// 解析成功的行数。
	ParsedSuccessNumber *string `json:"parsed_success_number,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowDbObjectTemplateProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectTemplateProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowDbObjectTemplateProgressResponse", string(data)}, " ")
}
