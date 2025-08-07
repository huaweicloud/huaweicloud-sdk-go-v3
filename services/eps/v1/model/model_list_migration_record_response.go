package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationRecordResponse Response Object
type ListMigrationRecordResponse struct {

	// 资源迁移记录列表
	Records        *[]ResourceMigrateRecord `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListMigrationRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationRecordResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationRecordResponse", string(data)}, " ")
}
