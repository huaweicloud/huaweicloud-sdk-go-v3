package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditDatabaseNewResponse Response Object
type AddAuditDatabaseNewResponse struct {

	// 数据库ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddAuditDatabaseNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditDatabaseNewResponse struct{}"
	}

	return strings.Join([]string{"AddAuditDatabaseNewResponse", string(data)}, " ")
}
