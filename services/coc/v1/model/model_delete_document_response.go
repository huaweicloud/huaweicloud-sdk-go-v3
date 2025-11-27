package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDocumentResponse Response Object
type DeleteDocumentResponse struct {

	// **参数解释：** 删除作业，系统返回的作业id。 **取值范围：** 不涉及。
	Data *string `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDocumentResponse struct{}"
	}

	return strings.Join([]string{"DeleteDocumentResponse", string(data)}, " ")
}
