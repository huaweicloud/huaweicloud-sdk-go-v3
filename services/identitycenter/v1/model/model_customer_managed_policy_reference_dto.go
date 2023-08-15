package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomerManagedPolicyReferenceDto 指定客户管理策略的名称和路径
type CustomerManagedPolicyReferenceDto struct {

	// IAM策略名称.
	Name string `json:"name"`

	// IAM策略路径，默认值/.
	Path *string `json:"path,omitempty"`
}

func (o CustomerManagedPolicyReferenceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerManagedPolicyReferenceDto struct{}"
	}

	return strings.Join([]string{"CustomerManagedPolicyReferenceDto", string(data)}, " ")
}
