package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetInstanceHealthReportConfigReq struct {

	// 健康报告模板ID。
	TemplateId string `json:"template_id"`
}

func (o SetInstanceHealthReportConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceHealthReportConfigReq struct{}"
	}

	return strings.Join([]string{"SetInstanceHealthReportConfigReq", string(data)}, " ")
}
