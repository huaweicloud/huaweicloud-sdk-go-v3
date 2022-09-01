package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 审计日志信息。
type Auditlog struct {

	// 审计日志ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 审计日志文件名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 审计日志大小，单位：KB。
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 审计日志开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 审计日志结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o Auditlog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Auditlog struct{}"
	}

	return strings.Join([]string{"Auditlog", string(data)}, " ")
}
