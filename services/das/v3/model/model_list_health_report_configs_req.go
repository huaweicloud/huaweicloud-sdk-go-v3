package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListHealthReportConfigsReq struct {

	// 实例ID列表，最大100条。
	InstanceIds []string `json:"instance_ids"`
}

func (o ListHealthReportConfigsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportConfigsReq struct{}"
	}

	return strings.Join([]string{"ListHealthReportConfigsReq", string(data)}, " ")
}
