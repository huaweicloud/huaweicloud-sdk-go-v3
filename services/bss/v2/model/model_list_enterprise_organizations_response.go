package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseOrganizationsResponse Response Object
type ListEnterpriseOrganizationsResponse struct {

	// 根节点ID，如果请求有parent_id，则该参数无值。
	RootId *string `json:"root_id,omitempty"`

	// 根节点名称，如果请求有parent_id，则该参数无值。  说明： 组织根节点没有设置组织名称时，可能为空。
	RootName *string `json:"root_name,omitempty"`

	// 子节点列表。
	ChildNodes     *[]EmChildNodeV2 `json:"child_nodes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEnterpriseOrganizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseOrganizationsResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseOrganizationsResponse", string(data)}, " ")
}
