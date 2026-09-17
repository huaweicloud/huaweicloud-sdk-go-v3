package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIssuesV2Response Response Object
type BatchCreateIssuesV2Response struct {

	// **参数解释**： 批量编辑的结果。
	Result *[]IssueEntity `json:"result,omitempty"`

	// **参数解释**： 返回状态。 **取值范围**： - success：响应成功 - error：响应失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 操作失败原因。 **取值范围**： 不涉及
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateIssuesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIssuesV2Response struct{}"
	}

	return strings.Join([]string{"BatchCreateIssuesV2Response", string(data)}, " ")
}
