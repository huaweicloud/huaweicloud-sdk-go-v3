package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiPolicyGroupsRequest Request Object
type DeleteAiPolicyGroupsRequest struct {
	Body *DeleteAiPolicyGroupsRequestInfo `json:"body,omitempty"`
}

func (o DeleteAiPolicyGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiPolicyGroupsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAiPolicyGroupsRequest", string(data)}, " ")
}
