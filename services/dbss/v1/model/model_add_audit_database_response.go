package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAuditDatabaseResponse Response Object
type AddAuditDatabaseResponse struct {

	// 数据库ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddAuditDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAuditDatabaseResponse struct{}"
	}

	return strings.Join([]string{"AddAuditDatabaseResponse", string(data)}, " ")
}
