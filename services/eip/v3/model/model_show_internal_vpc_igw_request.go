package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInternalVpcIgwRequest Request Object
type ShowInternalVpcIgwRequest struct {

	// 形式为\\\"fields=id&fields=project_id&...\\\"，支持字段：id/project_id/vpc_id/created_at/updated_at/igw_cluster
	Fields *[]string `json:"fields,omitempty"`

	// 虚拟igw的uuid
	VpcIgwId string `json:"vpc_igw_id"`
}

func (o ShowInternalVpcIgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternalVpcIgwRequest struct{}"
	}

	return strings.Join([]string{"ShowInternalVpcIgwRequest", string(data)}, " ")
}
