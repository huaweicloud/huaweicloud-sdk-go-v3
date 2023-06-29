package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubEnterpriseAccountRequest Request Object
type CreateSubEnterpriseAccountRequest struct {
	Body *CreateSubCustomerReqV2 `json:"body,omitempty"`
}

func (o CreateSubEnterpriseAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubEnterpriseAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateSubEnterpriseAccountRequest", string(data)}, " ")
}
