package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetRequest Request Object
type ListInternetRequest struct {

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// NAT网关名称。
	NatName *string `json:"nat_name,omitempty"`

	// EIP地址。
	Eip *string `json:"eip,omitempty"`
}

func (o ListInternetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetRequest struct{}"
	}

	return strings.Join([]string{"ListInternetRequest", string(data)}, " ")
}
