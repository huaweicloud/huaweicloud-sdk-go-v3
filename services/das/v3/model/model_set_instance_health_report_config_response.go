package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceHealthReportConfigResponse Response Object
type SetInstanceHealthReportConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetInstanceHealthReportConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceHealthReportConfigResponse struct{}"
	}

	return strings.Join([]string{"SetInstanceHealthReportConfigResponse", string(data)}, " ")
}
