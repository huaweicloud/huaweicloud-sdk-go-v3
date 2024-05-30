package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApproverResultData data，统一的返回结果的最外层数据结构。
type CreateApproverResultData struct {
	Value *ApproverVo `json:"value,omitempty"`
}

func (o CreateApproverResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApproverResultData struct{}"
	}

	return strings.Join([]string{"CreateApproverResultData", string(data)}, " ")
}
