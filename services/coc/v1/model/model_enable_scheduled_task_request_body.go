package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableScheduledTaskRequestBody struct {

	// 选择的四号提权单信息
	TicketInfos *[]TicketInfo `json:"ticket_infos,omitempty"`
}

func (o EnableScheduledTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableScheduledTaskRequestBody struct{}"
	}

	return strings.Join([]string{"EnableScheduledTaskRequestBody", string(data)}, " ")
}
