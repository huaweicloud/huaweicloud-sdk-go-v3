package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDocumentResponse Response Object
type ExecuteDocumentResponse struct {

	// **参数解释：**  执行作业，系统返回的作业工单id。 **取值范围：** 不涉及。
	Data *string `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentResponse", string(data)}, " ")
}
