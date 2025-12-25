package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserTicketRequest Request Object
type ShowUserTicketRequest struct {
}

func (o ShowUserTicketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserTicketRequest struct{}"
	}

	return strings.Join([]string{"ShowUserTicketRequest", string(data)}, " ")
}
