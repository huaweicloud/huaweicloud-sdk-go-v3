package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetConsistencyResultRequestBody 设置一致性校验的结果body体
type SetConsistencyResultRequestBody struct {

	// 一致性校验结果
	ConsistencyResult *[]ConsistencyResult `json:"consistency_result,omitempty"`
}

func (o SetConsistencyResultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetConsistencyResultRequestBody struct{}"
	}

	return strings.Join([]string{"SetConsistencyResultRequestBody", string(data)}, " ")
}
