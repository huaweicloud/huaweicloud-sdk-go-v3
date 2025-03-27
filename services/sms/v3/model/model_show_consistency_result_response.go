package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsistencyResultResponse Response Object
type ShowConsistencyResultResponse struct {

	// 一致性校验结果列表
	ResultList *[]ConsistencyResultRequestBodyResultList `json:"result_list,omitempty"`

	// 任务id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConsistencyResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsistencyResultResponse struct{}"
	}

	return strings.Join([]string{"ShowConsistencyResultResponse", string(data)}, " ")
}
