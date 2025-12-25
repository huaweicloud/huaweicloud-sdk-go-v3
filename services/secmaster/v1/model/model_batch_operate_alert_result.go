package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateAlertResult 批量操作告警返回对象
type BatchOperateAlertResult struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功id
	SuccessIds *[]string `json:"success_ids,omitempty"`
}

func (o BatchOperateAlertResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateAlertResult struct{}"
	}

	return strings.Join([]string{"BatchOperateAlertResult", string(data)}, " ")
}
