package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmApprovalsResultData 审批单撤回的返回结果，成功的个数。
type ConfirmApprovalsResultData struct {
	Value *BatchOperationVo `json:"value,omitempty"`
}

func (o ConfirmApprovalsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApprovalsResultData struct{}"
	}

	return strings.Join([]string{"ConfirmApprovalsResultData", string(data)}, " ")
}
