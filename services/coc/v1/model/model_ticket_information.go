package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TicketInformation struct {

	// 选择的四号提权单信息
	TicketInfos *[]TicketInfo `json:"ticket_infos,omitempty"`
}

func (o TicketInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TicketInformation struct{}"
	}

	return strings.Join([]string{"TicketInformation", string(data)}, " ")
}
