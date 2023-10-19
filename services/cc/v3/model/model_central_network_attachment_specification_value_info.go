package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkAttachmentSpecificationValueInfo 附件的额外信息。
type CentralNetworkAttachmentSpecificationValueInfo struct {

	// 资源ID标识符。
	EnterpriseRouterTableId *string `json:"enterprise_router_table_id,omitempty"`

	// 资源ID标识符。
	AttachedErId *string `json:"attached_er_id,omitempty"`

	ApprovedState *ApprovedStateEnum `json:"approved_state,omitempty"`

	HostedCloud *HostedCloudEnum `json:"hosted_cloud,omitempty"`

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
