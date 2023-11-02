package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertRequestBody 创建告警请求body体
type CreateAlertRequestBody struct {
	DataObject *Alert `json:"data_object"`
}

func (o CreateAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlertRequestBody", string(data)}, " ")
}
