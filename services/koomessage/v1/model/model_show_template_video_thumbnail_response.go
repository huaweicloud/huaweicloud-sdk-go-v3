package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateVideoThumbnailResponse Response Object
type ShowTemplateVideoThumbnailResponse struct {

	// 请求状态，固定200。
	Status *string `json:"status,omitempty"`

	// 状态描述。
	Message *string `json:"message,omitempty"`

	// 视频封面图列表。
	Data           *[]ListThumbnail `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTemplateVideoThumbnailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateVideoThumbnailResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateVideoThumbnailResponse", string(data)}, " ")
}
