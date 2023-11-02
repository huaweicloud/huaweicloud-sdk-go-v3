package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchOrderAlertsRequest Request Object
type CreateBatchOrderAlertsRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *OrderAlert `json:"body,omitempty"`
}

func (o CreateBatchOrderAlertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchOrderAlertsRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchOrderAlertsRequest", string(data)}, " ")
}
