package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableVacuumInfoOpen struct {

	// **参数解释**： 运行中的表信息。 **默认取值**： 不涉及
	VacuumRunningInfo *[]TableInfoOpen `json:"vacuum_running_info,omitempty"`

	// **参数解释**： 等待中的表信息。 **默认取值**： 不涉及
	VacuumWaitingInfo *[]TableInfoOpen `json:"vacuum_waiting_info,omitempty"`

	// **参数解释**： 已结束的表信息。 **默认取值**： 不涉及
	VacuumFinishedInfo *[]TableInfoOpen `json:"vacuum_finished_info,omitempty"`

	// **参数解释**： 取消的表信息。 **默认取值**： 不涉及
	VacuumCanceledInfo *[]TableInfoOpen `json:"vacuum_canceled_info,omitempty"`
}

func (o TableVacuumInfoOpen) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableVacuumInfoOpen struct{}"
	}

	return strings.Join([]string{"TableVacuumInfoOpen", string(data)}, " ")
}
