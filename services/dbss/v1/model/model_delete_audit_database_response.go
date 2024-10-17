package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditDatabaseResponse Response Object
type DeleteAuditDatabaseResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditDatabaseResponse", string(data)}, " ")
}
