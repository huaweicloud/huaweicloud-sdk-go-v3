package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsistencyResultRequestBody 一致性校验的结果body体
type ConsistencyResultRequestBody struct {

	// 校验结果
	ConsistencyResult []ConsistencyResult `json:"consistency_result"`

	// 检验完成时间
	FinishedTime int64 `json:"finished_time"`
}

func (o ConsistencyResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsistencyResultRequestBody struct{}"
	}

	return strings.Join([]string{"ConsistencyResultRequestBody", string(data)}, " ")
}
