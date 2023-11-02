package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlertRequestBody 更新告警请求body体
type ChangeAlertRequestBody struct {

	// 更新告警的ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`

	DataObject *Alert `json:"data_object,omitempty"`
}

func (o ChangeAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeAlertRequestBody", string(data)}, " ")
}
