package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditLogResponse Response Object
type ShowAuditLogResponse struct {

	// 全量SQL开关状态。 取值： - ON，表示开启 - OFF，表示关闭
	SwitchStatus   *string `json:"switch_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditLogResponse", string(data)}, " ")
}
