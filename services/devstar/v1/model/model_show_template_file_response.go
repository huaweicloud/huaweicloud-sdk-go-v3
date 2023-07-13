package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateFileResponse Response Object
type ShowTemplateFileResponse struct {

	// 文件内容（返回的文件内容为encoding指定的编码格式编码后的内容）。
	Content *string `json:"content,omitempty"`

	// 内容编码格式(固定base64)。
	Encoding *string `json:"encoding,omitempty"`

	// 文件名。
	FileName *string `json:"file_name,omitempty"`

	// 文件相对路径。
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型。
	FileType       *string `json:"file_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateFileResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateFileResponse", string(data)}, " ")
}
