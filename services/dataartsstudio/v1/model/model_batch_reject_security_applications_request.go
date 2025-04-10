package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRejectSecurityApplicationsRequest Request Object
type BatchRejectSecurityApplicationsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchApproveRequestBody `json:"body,omitempty"`
}

func (o BatchRejectSecurityApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRejectSecurityApplicationsRequest struct{}"
	}

	return strings.Join([]string{"BatchRejectSecurityApplicationsRequest", string(data)}, " ")
}
