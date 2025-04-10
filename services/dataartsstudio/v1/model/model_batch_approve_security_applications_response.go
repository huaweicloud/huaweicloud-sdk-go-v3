package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchApproveSecurityApplicationsResponse Response Object
type BatchApproveSecurityApplicationsResponse struct {

	// 工单批量审批结果
	Body           *[]ApproveResult `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchApproveSecurityApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveSecurityApplicationsResponse struct{}"
	}

	return strings.Join([]string{"BatchApproveSecurityApplicationsResponse", string(data)}, " ")
}
