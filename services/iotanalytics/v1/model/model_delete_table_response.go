package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableResponse Response Object
type DeleteTableResponse struct {

	// 被删除表ID
	TableId *string `json:"table_id,omitempty"`

	// 被删除表名。
	TableName      *string `json:"table_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableResponse", string(data)}, " ")
}
