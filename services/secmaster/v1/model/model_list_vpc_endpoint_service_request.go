package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcEndpointServiceRequest Request Object
type ListVpcEndpointServiceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListVpcEndpointServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcEndpointServiceRequest struct{}"
	}

	return strings.Join([]string{"ListVpcEndpointServiceRequest", string(data)}, " ")
}
