package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventModel 事件或者告警元数据。
type EventModel struct {

	// 指定上报的事件或者告警产生的时间。仅支持UTC毫秒级时间戳。  例如：2024-10-16 16:03:01需要通过工具转换成UTC毫秒级时间戳：1702759381000  当action值为空时，即上报事件或告警时需要时指定该参数。
	StartsAt *int64 `json:"starts_at,omitempty"`

	// 指定清除的事件或者告警清除的时间。仅支持UTC毫秒级时间戳。默认值为0，表示没有清除告警。  例如：2024-10-16 16:03:01需要通过工具转换成UTC毫秒级时间戳：1702759381000  当action值为clear时，即清除告警时需要时指定该参数。
	EndsAt *int64 `json:"ends_at,omitempty"`

	// 指定AOM自动清除超期告警的时间，最长清除时间不超过15天。单位：毫秒数，一分钟则填写为60000。例如该时间设置为5天的告警，对应毫秒数：7200 * 60000（即：5天 * 24小时 * 60分钟 * 60000毫秒）。如果上报告警时没指定该时间，则默认清除时间为15天。 当action值为clear时，即清除告警时不需要指定该参数。
	Timeout *int64 `json:"timeout,omitempty"`

	// 待上报的事件或者告警的详细信息，为key:value键值对形式。支持如下必填字段： - event_name：事件或者告警名称，类型为String； - event_severity：事件或告警级别。类型为String，支持四种级别：    - Critical：紧急    - Major：重要    - Minor：次要    - Info：提示 - event_type：事件或告警类别。类型为String，支持两种类别：   - event：告警事件   - alarm：普通告警 - resource_provider：事件对应云服务名称。类型为String；  - resource_type：事件对应资源类型。类型为String；  - resource_id：事件对应资源信息。类型为String。 metadata中的value长度为1到2048字符串。
	Metadata map[string]string `json:"metadata"`

	// 事件或者告警附加字段，可以为空。
	Annotations map[string]interface{} `json:"annotations,omitempty"`

	// 事件或者告警预留字段，可以为空。
	AttachRule map[string]interface{} `json:"attach_rule,omitempty"`

	// 事件或者告警id，产生事件或告警时，系统会自动生成。  当action值为clear时，即清除告警时需要时指定该参数。上报事件或告警时无需传入该参数。
	Id *string `json:"id,omitempty"`
}

func (o EventModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventModel struct{}"
	}

	return strings.Join([]string{"EventModel", string(data)}, " ")
}
