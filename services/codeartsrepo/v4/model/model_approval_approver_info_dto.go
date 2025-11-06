package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalApproverInfoDto struct {

	// **参数解释：** 需要更新的审核人ID列表 数字以英文逗号分隔 111,222 **取值范围：** 不涉及。
	ApproverIds *string `json:"approver_ids,omitempty"`
}

func (o ApprovalApproverInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalApproverInfoDto struct{}"
	}

	return strings.Join([]string{"ApprovalApproverInfoDto", string(data)}, " ")
}
