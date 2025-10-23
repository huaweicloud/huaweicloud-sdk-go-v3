package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportSettingResponse Response Object
type UpdateReportSettingResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReportSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateReportSettingResponse", string(data)}, " ")
}
