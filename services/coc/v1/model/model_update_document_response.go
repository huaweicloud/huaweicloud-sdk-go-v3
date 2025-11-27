package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDocumentResponse Response Object
type UpdateDocumentResponse struct {

	// **参数解释：** 编辑作业，系统返回的作业id。 **取值范围：** 不涉及。
	Data *string `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocumentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDocumentResponse", string(data)}, " ")
}
