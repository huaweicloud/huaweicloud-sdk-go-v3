package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadExtensionFileRequest struct {

	// 插件类型。目前只支持CloudIDE
	Official string `json:"official"`

	Body *UploadExtensionFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadExtensionFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadExtensionFileRequest struct{}"
	}

	return strings.Join([]string{"UploadExtensionFileRequest", string(data)}, " ")
}
