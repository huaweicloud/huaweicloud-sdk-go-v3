package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackUpInfoResponse Response Object
type ShowBackUpInfoResponse struct {

	// 是否备份
	IsAdditionalBackup *bool `json:"is_additional_backup,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 节点IP
	NodeIp *string `json:"node_ip,omitempty"`

	// 节点角色
	NodeRole *string `json:"node_role,omitempty"`

	// 备份周期
	BackupPeriod   *string `json:"backup_period,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackUpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackUpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBackUpInfoResponse", string(data)}, " ")
}
