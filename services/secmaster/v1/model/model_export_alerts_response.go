package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAlertsResponse Response Object
type ExportAlertsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAlertsResponse struct{}"
	}

	return strings.Join([]string{"ExportAlertsResponse", string(data)}, " ")
}
