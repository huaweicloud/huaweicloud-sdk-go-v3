package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchApproveApplyRequest Request Object
type BatchApproveApplyRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApplyIdsForApproveApply `json:"body,omitempty"`
}

func (o BatchApproveApplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveApplyRequest struct{}"
	}

	return strings.Join([]string{"BatchApproveApplyRequest", string(data)}, " ")
}
