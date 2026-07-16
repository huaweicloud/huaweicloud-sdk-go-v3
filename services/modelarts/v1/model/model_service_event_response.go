package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceEventResponse 服务事件信息响应数据模型
type ServiceEventResponse struct {

	// **参数解释：** 服务事件ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务ID。 **取值范围：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 服务版本ID。 **取值范围：** 不涉及。
	ServiceVersionId *string `json:"service_version_id,omitempty"`

	// **参数解释：** 服务事件发生计数 **取值范围：** 不涉及。
	EventCount *int32 `json:"event_count,omitempty"`

	// **参数解释：** 服务事件类型：NORMAL/ABNORMAL/WARNING **取值范围：** 不涉及。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释：** 服务事件信息（英文） **取值范围：** 不涉及。
	EventInfo *string `json:"event_info,omitempty"`

	// **参数解释：** 服务事件信息（中文） **取值范围：** 不涉及。
	EventInfoCn *string `json:"event_info_cn,omitempty"`

	// **参数解释：** 服务事件第一次发生时间 **取值范围：** 不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释：** 服务事件最后发生时间 **取值范围：** 不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`
}

func (o ServiceEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceEventResponse struct{}"
	}

	return strings.Join([]string{"ServiceEventResponse", string(data)}, " ")
}
