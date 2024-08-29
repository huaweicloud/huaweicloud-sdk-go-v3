package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantVpcIgwsRequest Request Object
type ListTenantVpcIgwsRequest struct {

	// 形式为\\\"fields=id&fields=project_id&...\\\"，支持字段：id/project_id/vpc_id/created_at/updated_at/name
	Fields *[]string `json:"fields,omitempty"`

	// 虚拟IGW的uuid
	Id *string `json:"id,omitempty"`

	// 虚拟igw所在的vpcid
	VpcId *string `json:"vpc_id,omitempty"`

	// 虚拟igw的名称
	Name *string `json:"name,omitempty"`

	// 排序，形式为\"sort_key=i2a_id&sort_dir=asc\"  支持字段：id/created_at/updated_at
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向  取值范围：asc、desc
	SortDir *string `json:"sort_dir,omitempty"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`
}

func (o ListTenantVpcIgwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantVpcIgwsRequest struct{}"
	}

	return strings.Join([]string{"ListTenantVpcIgwsRequest", string(data)}, " ")
}
