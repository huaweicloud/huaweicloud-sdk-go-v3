package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryPublicationInfo 发布详情。
type QueryPublicationInfo struct {

	// 发布id。
	Id *string `json:"id,omitempty"`

	// 发布状态。normal，abnormal，creating，modifying，createfail
	Status *string `json:"status,omitempty"`

	// 发布名。
	PublicationName *string `json:"publication_name,omitempty"`

	// 发布数据库名。
	PublicationDatabase *string `json:"publication_database,omitempty"`

	// 订阅数。
	SubscriptionCount *int32 `json:"subscription_count,omitempty"`

	SubscriptionOptions *SubscriptionOptionsResult `json:"subscription_options,omitempty"`

	JobSchedule *UsedJobSchedule `json:"job_schedule,omitempty"`

	// 是否选择所有数据表。
	IsSelectAllTable *bool `json:"is_select_all_table,omitempty"`

	// 全选所有表后需要去除的表。
	ExtendTables *[]string `json:"extend_tables,omitempty"`

	// 发布数据表。
	Tables *[]PublicationTableInfoResponse `json:"tables,omitempty"`
}

func (o QueryPublicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryPublicationInfo struct{}"
	}

	return strings.Join([]string{"QueryPublicationInfo", string(data)}, " ")
}
