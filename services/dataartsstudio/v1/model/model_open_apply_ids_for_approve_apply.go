package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenApplyIdsForApproveApply struct {

	// 申请结果
	ApplyResult *bool `json:"apply_result,omitempty"`

	// 申请编号列表
	ApplyIds *[]string `json:"apply_ids,omitempty"`
}

func (o OpenApplyIdsForApproveApply) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApplyIdsForApproveApply struct{}"
	}

	return strings.Join([]string{"OpenApplyIdsForApproveApply", string(data)}, " ")
}
