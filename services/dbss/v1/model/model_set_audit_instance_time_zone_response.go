package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditInstanceTimeZoneResponse Response Object
type SetAuditInstanceTimeZoneResponse struct {

	// 状态  - success：成功  - fail：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditInstanceTimeZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditInstanceTimeZoneResponse struct{}"
	}

	return strings.Join([]string{"SetAuditInstanceTimeZoneResponse", string(data)}, " ")
}
