package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcInfo vpc信息。
type VpcInfo struct {

	// vpc status。
	Status *string `json:"status,omitempty"`

	// vpcId。
	Id *string `json:"id,omitempty"`

	// vpc 名。
	Name *string `json:"name,omitempty"`

	// vpc创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// vpc可用子网的范围。
	Cidr *string `json:"cidr,omitempty"`

	// 租户Id。
	TenantId *string `json:"tenant_id,omitempty"`
}

func (o VpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcInfo struct{}"
	}

	return strings.Join([]string{"VpcInfo", string(data)}, " ")
}
