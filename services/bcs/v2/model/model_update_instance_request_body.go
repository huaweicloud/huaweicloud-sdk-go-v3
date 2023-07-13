package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRequestBody 添加节点和添加组织，添加组织是需要提供pvc_name
type UpdateInstanceRequestBody struct {

	// 添加节点的组织列表
	NodeOrgs []NodeOrgs `json:"node_orgs"`

	// ief添加组织时，ief节点信息。绑定模式的IEF服务，新增组织时，该字段必填
	Publicips *[]IefNodeinfo `json:"publicips,omitempty"`

	// 是否是删除组织
	IsDeleteOrg *bool `json:"is_delete_org,omitempty"`
}

func (o UpdateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequestBody", string(data)}, " ")
}
