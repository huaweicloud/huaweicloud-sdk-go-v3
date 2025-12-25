package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeAlertRequestBody 批量更新告警请求body体
type BatchChangeAlertRequestBody struct {

	// 更新告警的ID列表
	BatchIds []string `json:"batch_ids"`

	DataObject *Alert `json:"data_object"`
}

func (o BatchChangeAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeAlertRequestBody struct{}"
	}

	return strings.Join([]string{"BatchChangeAlertRequestBody", string(data)}, " ")
}
