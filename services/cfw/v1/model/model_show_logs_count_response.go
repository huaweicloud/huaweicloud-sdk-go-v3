package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogsCountResponse Response Object
type ShowLogsCountResponse struct {
	Data           *ShowLogsCountRespData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowLogsCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsCountResponse struct{}"
	}

	return strings.Join([]string{"ShowLogsCountResponse", string(data)}, " ")
}
