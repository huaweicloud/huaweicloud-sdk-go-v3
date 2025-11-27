package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicationsRequestBody 创建发布详情。
type CreatePublicationsRequestBody struct {

	// 发布名称。  发布名称长度可在5～64个字符之间，由字母、数字或下划线组成，不能包含其他特殊字符。
	PublicationName string `json:"publication_name"`

	// 发布数据库名。
	PublicationDatabase string `json:"publication_database"`

	// 是否立即创建快照。
	IsCreateSnapshotImmediately string `json:"is_create_snapshot_immediately"`

	SubscriptionOptions *SubscriptionOption `json:"subscription_options,omitempty"`

	JobSchedule *OperateUsedJobSchedule `json:"job_schedule"`

	// 是否选择所有数据表。
	IsSelectAllTable *bool `json:"is_select_all_table,omitempty"`

	// 全选所有表后需要去除的表。
	ExtendTables *[]string `json:"extend_tables,omitempty"`

	// 发布数据表。
	Tables []PublicationTableInfoRequest `json:"tables"`
}

func (o CreatePublicationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicationsRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePublicationsRequestBody", string(data)}, " ")
}
