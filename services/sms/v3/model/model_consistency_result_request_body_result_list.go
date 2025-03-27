package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsistencyResultRequestBodyResultList struct {

	// 校验完成时间
	FinishedTime *int64 `json:"finished_time,omitempty"`

	// 校验结果
	CheckResult *string `json:"check_result,omitempty"`

	// 校验结果
	ConsistencyResult *[]ConsistencyResult `json:"consistency_result,omitempty"`
}

func (o ConsistencyResultRequestBodyResultList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsistencyResultRequestBodyResultList struct{}"
	}

	return strings.Join([]string{"ConsistencyResultRequestBodyResultList", string(data)}, " ")
}
