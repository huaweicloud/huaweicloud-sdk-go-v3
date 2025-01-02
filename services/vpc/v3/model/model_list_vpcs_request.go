package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcsRequest Request Object
type ListVpcsRequest struct {

	// 功能说明：每页返回的个数 取值范围：0-2000
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// VPC资源ID。可以使用该字段过滤VPC
	Id *[]string `json:"id,omitempty"`

	// 功能说明：企业项目ID。可以使用该字段过滤某个企业项目下的VPC。 取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。若需要查询当前用户所有企业项目绑定的VPC，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// VPC的name信息，可以使用该字段过滤VPC
	Name *[]string `json:"name,omitempty"`

	// VPC的描述信息。可以使用该字段过滤VPC
	Description *[]string `json:"description,omitempty"`

	// VPC的CIDR。可以使用该字段过滤VPC
	Cidr *[]string `json:"cidr,omitempty"`
}

func (o ListVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsRequest", string(data)}, " ")
}
