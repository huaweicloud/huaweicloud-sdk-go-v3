package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseRequest Request Object
type DeleteDatabaseRequest struct {

	// 删除的数据库名称。
	DatabaseName string `json:"database_name"`

	Async *bool `json:"async,omitempty"`

	Cascade *bool `json:"cascade,omitempty"`
}

func (o DeleteDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}
