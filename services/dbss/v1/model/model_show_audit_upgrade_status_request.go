package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditUpgradeStatusRequest Request Object
type ShowAuditUpgradeStatusRequest struct {
}

func (o ShowAuditUpgradeStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditUpgradeStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditUpgradeStatusRequest", string(data)}, " ")
}
