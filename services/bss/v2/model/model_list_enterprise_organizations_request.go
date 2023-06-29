package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseOrganizationsRequest Request Object
type ListEnterpriseOrganizationsRequest struct {

	// 是否递归查询。0：不递归（默认）1：递归如果不递归，只返回起始节点的直接子节点。此参数不携带或携带值为空时，不作为筛选条件。
	RecursiveQuery *int32 `json:"recursive_query,omitempty"`

	// 指定的节点ID。为空则从根节点查起。此参数不携带或携带值为空时，不作为筛选条件。 说明： 此参数须由纯数字组成。
	ParentId *string `json:"parent_id,omitempty"`
}

func (o ListEnterpriseOrganizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseOrganizationsRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseOrganizationsRequest", string(data)}, " ")
}
