package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResultVoIssueWithReasonVo 批量操作返回体。
type BatchResultVoIssueWithReasonVo struct {

	// **参数解释**： 批量操作成功条数。 **取值范围**： 不涉及。
	SuccessNum *int32 `json:"success_num,omitempty"`

	// **参数解释**： 批量操作失败条数。 **取值范围**： 不涉及。
	FailNum *int32 `json:"fail_num,omitempty"`

	// **参数解释**： 批量操作失败数据及失败原因。 **取值范围**： 不涉及。
	Failed *[]IssueWithReasonVo `json:"failed,omitempty"`
}

func (o BatchResultVoIssueWithReasonVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResultVoIssueWithReasonVo struct{}"
	}

	return strings.Join([]string{"BatchResultVoIssueWithReasonVo", string(data)}, " ")
}
