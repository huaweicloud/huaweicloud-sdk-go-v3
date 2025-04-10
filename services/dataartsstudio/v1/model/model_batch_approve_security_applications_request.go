package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchApproveSecurityApplicationsRequest Request Object
type BatchApproveSecurityApplicationsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchApproveRequestBody `json:"body,omitempty"`
}

func (o BatchApproveSecurityApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchApproveSecurityApplicationsRequest struct{}"
	}

	return strings.Join([]string{"BatchApproveSecurityApplicationsRequest", string(data)}, " ")
}
