package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlobsResponse Response Object
type ShowBlobsResponse struct {

	// **参数解释：** 文件字节大小。 **约束限制：** 不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释：** 文件编码方式。 **约束限制：** - base64。
	Encoding *string `json:"encoding,omitempty"`

	// **参数解释：** 经过base64编码后的文件内容。 **约束限制：** 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释：** blob文件ID。 **约束限制：** 不涉及。
	BlobId *string `json:"blob_id,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBlobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlobsResponse struct{}"
	}

	return strings.Join([]string{"ShowBlobsResponse", string(data)}, " ")
}
