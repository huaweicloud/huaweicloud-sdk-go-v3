package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuoteDatabaseResultRsp struct {

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 源数据库id
	SourceDatabaseId *string `json:"source_database_id,omitempty"`

	// 引用到项目后的数据库id
	DestinationDatabaseId *string `json:"destination_database_id,omitempty"`

	// 引用到项目后的数据库名称
	DestinationDatabaseName *string `json:"destination_database_name,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 导入结果
	Status *string `json:"status,omitempty"`
}

func (o QuoteDatabaseResultRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuoteDatabaseResultRsp struct{}"
	}

	return strings.Join([]string{"QuoteDatabaseResultRsp", string(data)}, " ")
}
