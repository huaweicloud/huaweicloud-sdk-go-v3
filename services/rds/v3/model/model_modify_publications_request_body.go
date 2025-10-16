package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPublicationsRequestBody 修改发布详情。
type ModifyPublicationsRequestBody struct {
	SubscriptionOptions *SubscriptionOption `json:"subscription_options,omitempty"`

	JobSchedule *OperateUsedJobSchedule `json:"job_schedule,omitempty"`

	// 是否选择所有数据表。
	IsSelectAllTable *bool `json:"is_select_all_table,omitempty"`

	// 全选所有表后需要去除的表。
	ExtendTables *[]string `json:"extend_tables,omitempty"`

	// 发布数据表。
	Tables *[]PublicationTableInfoRequest `json:"tables,omitempty"`
}

func (o ModifyPublicationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicationsRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyPublicationsRequestBody", string(data)}, " ")
}
