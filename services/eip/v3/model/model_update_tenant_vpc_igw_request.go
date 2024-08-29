package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantVpcIgwRequest Request Object
type UpdateTenantVpcIgwRequest struct {

	// 形式为\\\"fields=id&fields=project_id&...\\\"，支持字段：id/project_id/vpc_id/created_at/updated_at/name
	Fields *[]string `json:"fields,omitempty"`

	// vpc-igw的uuid
	VpcIgwId string `json:"vpc_igw_id"`

	Body *UpdateTenantVpcIgwRequestBody `json:"body,omitempty"`
}

func (o UpdateTenantVpcIgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantVpcIgwRequest struct{}"
	}

	return strings.Join([]string{"UpdateTenantVpcIgwRequest", string(data)}, " ")
}
