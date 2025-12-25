package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlertRequestBody 更新告警请求body体
type ChangeAlertRequestBody struct {
	DataObject *Alert `json:"data_object"`
}

func (o ChangeAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeAlertRequestBody", string(data)}, " ")
}
