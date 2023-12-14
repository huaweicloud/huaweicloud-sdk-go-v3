package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsistencyResultResponse Response Object
type ShowConsistencyResultResponse struct {

	// 校验结果
	ConsistencyResult *[]ConsistencyResult `json:"consistency_result,omitempty"`

	// 检验完成时间
	FinishedTime   *int64 `json:"finished_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowConsistencyResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsistencyResultResponse struct{}"
	}

	return strings.Join([]string{"ShowConsistencyResultResponse", string(data)}, " ")
}
