package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTableResponse struct {

	// 被删除表ID
	TableId *string `json:"table_id,omitempty" xml:"table_id"`

	// 被删除表名。
	TableName      *string `json:"table_name,omitempty" xml:"table_name"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableResponse", string(data)}, " ")
}
