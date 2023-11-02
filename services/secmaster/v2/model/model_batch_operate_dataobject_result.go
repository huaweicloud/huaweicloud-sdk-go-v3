package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateDataobjectResult 批量操作告警返回对象
type BatchOperateDataobjectResult struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功id
	SuccessIds *[]string `json:"success_ids,omitempty"`
}

func (o BatchOperateDataobjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateDataobjectResult struct{}"
	}

	return strings.Join([]string{"BatchOperateDataobjectResult", string(data)}, " ")
}
