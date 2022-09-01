package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTemplateFileResponse struct {

	// 文件内容（返回的文件内容为encoding指定的编码格式编码后的内容）。
	Content *string `json:"content,omitempty" xml:"content"`

	// 内容编码格式(固定base64)。
	Encoding *string `json:"encoding,omitempty" xml:"encoding"`

	// 文件名。
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 文件相对路径。
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件类型。
	FileType       *string `json:"file_type,omitempty" xml:"file_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTemplateFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateFileResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateFileResponse", string(data)}, " ")
}
