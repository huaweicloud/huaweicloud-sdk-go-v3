package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportSettingResponse Response Object
type CreateReportSettingResponse struct {

	// 配置ID
	SettingId      *string `json:"setting_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReportSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportSettingResponse struct{}"
	}

	return strings.Join([]string{"CreateReportSettingResponse", string(data)}, " ")
}
