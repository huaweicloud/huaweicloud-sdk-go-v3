package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetHealthReportConfigReq struct {

	// 实例ID列表，最大100条。
	InstanceIds []string `json:"instance_ids"`

	// 健康报告模板ID。
	TemplateId string `json:"template_id"`
}

func (o BatchSetHealthReportConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetHealthReportConfigReq struct{}"
	}

	return strings.Join([]string{"BatchSetHealthReportConfigReq", string(data)}, " ")
}
