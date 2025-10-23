package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportSettingResponse Response Object
type DeleteReportSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteReportSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportSettingResponse struct{}"
	}

	return strings.Join([]string{"DeleteReportSettingResponse", string(data)}, " ")
}
