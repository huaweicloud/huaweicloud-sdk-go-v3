package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryOperationsResult 参数修改历史记录明细信息。
type ListHistoryOperationsResult struct {

	// 参数名称。
	ParameterName string `json:"parameter_name"`

	// 修改前参数值。
	OldValue string `json:"old_value"`

	// 修改后参数值。
	NewValue string `json:"new_value"`

	// 修改状态 (SUCCESS | FAILED)。
	UpdateResult string `json:"update_result"`

	// 修改时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	UpdatedAt string `json:"updated_at"`
}

func (o ListHistoryOperationsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOperationsResult struct{}"
	}

	return strings.Join([]string{"ListHistoryOperationsResult", string(data)}, " ")
}
