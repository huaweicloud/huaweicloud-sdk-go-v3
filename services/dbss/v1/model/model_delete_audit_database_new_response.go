package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditDatabaseNewResponse Response Object
type DeleteAuditDatabaseNewResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditDatabaseNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditDatabaseNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditDatabaseNewResponse", string(data)}, " ")
}
