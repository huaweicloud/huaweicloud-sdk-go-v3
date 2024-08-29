package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTenantVpcIgwRequest Request Object
type CreateTenantVpcIgwRequest struct {

	// 形式为\\\"fields=id&fields=project_id&...\\\"，支持字段：id/project_id/vpc_id/created_at/updated_at/name
	Fields *[]string `json:"fields,omitempty"`

	Body *CreateTenantVpcIgwRequestBody `json:"body,omitempty"`
}

func (o CreateTenantVpcIgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTenantVpcIgwRequest struct{}"
	}

	return strings.Join([]string{"CreateTenantVpcIgwRequest", string(data)}, " ")
}
