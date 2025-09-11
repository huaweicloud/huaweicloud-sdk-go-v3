package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditUpgradeStatusResponse Response Object
type ShowAuditUpgradeStatusResponse struct {

	// 实例升级信息
	AuditUpgradeInfos *[]AuditUpgradeStatus `json:"audit_upgrade_infos,omitempty"`
	HttpStatusCode    int                   `json:"-"`
}

func (o ShowAuditUpgradeStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditUpgradeStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditUpgradeStatusResponse", string(data)}, " ")
}
