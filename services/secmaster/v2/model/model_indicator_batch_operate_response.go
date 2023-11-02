package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorBatchOperateResponse 情报响应参数
type IndicatorBatchOperateResponse struct {

	// 成功ID列表
	SuccessIds []string `json:"success_ids"`

	// 失败ID列表
	ErrorIds *[]string `json:"error_ids,omitempty"`
}

func (o IndicatorBatchOperateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorBatchOperateResponse struct{}"
	}

	return strings.Join([]string{"IndicatorBatchOperateResponse", string(data)}, " ")
}
