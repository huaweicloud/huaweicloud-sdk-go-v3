package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoIncrementUsageResponse Response Object
type ListAutoIncrementUsageResponse struct {

	// 任务执行状态。 取值：  值为“Running”，表示任务正在执行。  值为“Completed”，表示任务执行成功。  值为“Failed”，表示任务执行失败。
	Status *string `json:"status,omitempty"`

	// 数量
	Total *int32 `json:"total,omitempty"`

	// 自增 ID 使用数据列表。
	AutoIncrementUsageList *[]AutoIncrementUsageRespAutoIncrementUsageList `json:"auto_increment_usage_list,omitempty"`
	HttpStatusCode         int                                             `json:"-"`
}

func (o ListAutoIncrementUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoIncrementUsageResponse struct{}"
	}

	return strings.Join([]string{"ListAutoIncrementUsageResponse", string(data)}, " ")
}
