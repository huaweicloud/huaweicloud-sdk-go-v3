package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIssueNoteResultResult **参数解释**： 删除操作返回的业务结果，包含删除状态标识。
type DeleteIssueNoteResultResult struct {

	// **参数解释**： 删除工作项评论的业务处理结果。 **取值范围**： - success：删除工作项评论成功。
	Status *string `json:"status,omitempty"`
}

func (o DeleteIssueNoteResultResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIssueNoteResultResult struct{}"
	}

	return strings.Join([]string{"DeleteIssueNoteResultResult", string(data)}, " ")
}
