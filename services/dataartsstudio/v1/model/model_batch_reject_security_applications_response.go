package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejectSecurityApplicationsResponse Response Object
type BatchRejectSecurityApplicationsResponse struct {

	// 工单批量审批结果
	Body           *[]ApproveResult `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchRejectSecurityApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejectSecurityApplicationsResponse struct{}"
	}

	return strings.Join([]string{"BatchRejectSecurityApplicationsResponse", string(data)}, " ")
}
