package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServiceEventsRequest Request Object
type ListInferServiceEventsRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`

	// **参数解释：** 在线服务事件类型。 **约束限制：** 不涉及。 **取值范围：** - NORMAL：正常 - ABNORMAL：异常 - WARNING：警告 **默认取值：** 不涉及。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释：** 事件开始时间。 **约束限制：** 不涉及。 **取值范围：** 毫秒级时间戳，13位数字。 **默认取值：** 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 事件结束时间。 **约束限制：** 不涉及。 **取值范围：** 毫秒级时间戳，13位数字。 **默认取值：** 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 事件信息过滤关键字。 **约束限制：** 不支持'\";%_*!@#$&\\这些字符的查询。
	EventInfoKey *string `json:"event_info_key,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** offset必须是limit的整数倍。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 排序字段，多个字段以\",\"分隔，支持create_at, update_at，默认值update_at。 **约束限制：** 不涉及。 **取值范围：** - create_at：按创建时间排序。 - update_at：按更新时间排序。 **默认取值：** update_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释：** 排序方式。 **约束限制：** 不涉及。 **取值范围：** - ASC: 递增排序。 - DESC: 递减排序。 **默认取值：** DESC。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListInferServiceEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServiceEventsRequest struct{}"
	}

	return strings.Join([]string{"ListInferServiceEventsRequest", string(data)}, " ")
}
