package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWdrSnapshotStatusResponse Response Object
type ShowWdrSnapshotStatusResponse struct {

	// **参数解释**： 开关状态。 **取值范围**： - ON：WDR快照开启。 - OFF：WDR快照关闭。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 快照生成的最早时间，格式为“yyyy-mm-dd hh:mm:ss timezone”。其中timezone是指指时区偏移量，例如北京时间偏移显示为+08。若内核未生成快照，则该值为空。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 快照生成的最新时间，格式为“yyyy-mm-dd hh:mm:ss timezone”。其中timezone是指指时区偏移量，例如北京时间偏移显示为+08。若内核未生成快照，则该值为空。 **取值范围**： 不涉及。
	EndTime        *string `json:"end_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWdrSnapshotStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWdrSnapshotStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowWdrSnapshotStatusResponse", string(data)}, " ")
}
