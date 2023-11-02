package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertRequestBody 删除告警请求body体
type DeleteAlertRequestBody struct {

	// 删除告警的ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DeleteAlertRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteAlertRequestBody", string(data)}, " ")
}
