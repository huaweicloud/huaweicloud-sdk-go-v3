package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReadmeFileResponse Response Object
type ShowReadmeFileResponse struct {

	// **参数解释：** blob文件ID。 **约束限制：** 不涉及。
	BlobId *string `json:"blob_id,omitempty"`

	// **参数解释：** 经过base64编码后的文件内容。 **约束限制：** 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释：** 文件编码方式。 **约束限制：** - base64。
	Encoding *interface{} `json:"encoding,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// **参数解释：** 文件字节大小。 **约束限制：** 不涉及。
	Size *int64 `json:"size,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReadmeFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReadmeFileResponse struct{}"
	}

	return strings.Join([]string{"ShowReadmeFileResponse", string(data)}, " ")
}
