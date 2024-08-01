package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProcessesAuditLogRequest Request Object
type ShowProcessesAuditLogRequest struct {

	// DDM实例ID或关联RDS的ID。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。取值必须为数字，且不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~128。不传该参数时，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间，UTC time，精确到毫秒。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime string `json:"start_time"`

	// 结束时间，UTC time，精确到毫秒。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。结束时间与开始时间，间隔不能超过7天。
	EndTime string `json:"end_time"`
}

func (o ShowProcessesAuditLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProcessesAuditLogRequest struct{}"
	}

	return strings.Join([]string{"ShowProcessesAuditLogRequest", string(data)}, " ")
}
