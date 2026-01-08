package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportPolicyGroupsRequest Request Object
type ExportPolicyGroupsRequest struct {
	Body *ExportPolicyGroupsReq `json:"body,omitempty"`
}

func (o ExportPolicyGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPolicyGroupsRequest struct{}"
	}

	return strings.Join([]string{"ExportPolicyGroupsRequest", string(data)}, " ")
}
