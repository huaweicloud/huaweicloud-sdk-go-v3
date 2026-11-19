package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BinlogInfo Binlog日志详情。
type BinlogInfo struct {

	// **参数解释**：  Binlog日志文件ID。   **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  Binlog日志文件名字。   **取值范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  Binlog日志对应的实例ID。   **取值范围**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  Binlog日志对应的实例名称。   **取值范围**：  不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**：  Binlog日志文件大小，单位是KB。   **取值范围**：  不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：  Binlog日志文件中上一轮任务的最后一个事务的提交时间。   **取值范围**：  不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  Binlog日志文件中最后一个事务的提交时间。   **取值范围**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o BinlogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogInfo struct{}"
	}

	return strings.Join([]string{"BinlogInfo", string(data)}, " ")
}
