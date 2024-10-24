package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkAttachmentSpecificationValueInfo 附件的额外信息。
type CentralNetworkAttachmentSpecificationValueInfo struct {

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`

	// 连接的父资源ID，这里表示企业路由器ID。
	AttachmentParentInstanceId *string `json:"attachment_parent_instance_id,omitempty"`

	HostedCloud *HostedCloudEnum `json:"hosted_cloud,omitempty"`

	ApprovedState *ApprovedStateEnum `json:"approved_state,omitempty"`

	// 审批拒绝创建附件的原因。
	Reason *string `json:"reason,omitempty"`
}

func (o CentralNetworkAttachmentSpecificationValueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkAttachmentSpecificationValueInfo struct{}"
	}

	return strings.Join([]string{"CentralNetworkAttachmentSpecificationValueInfo", string(data)}, " ")
}
