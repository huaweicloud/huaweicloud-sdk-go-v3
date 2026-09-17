package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateIpdIssuesResponse Response Object
type BatchUpdateIpdIssuesResponse struct {

	// **参数解释**： 批量编辑的结果。
	Result *[]IssueBatchOperateEntitiesResult `json:"result,omitempty"`

	// **参数解释**： 返回状态。 **取值范围**： - success：响应成功 - error：响应失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 操作失败原因。 **取值范围**： 不涉及
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateIpdIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIpdIssuesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateIpdIssuesResponse", string(data)}, " ")
}
