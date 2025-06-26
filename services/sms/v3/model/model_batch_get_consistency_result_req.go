package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetConsistencyResultReq 所有需要获取一致性校验结果任务的请求参数
type BatchGetConsistencyResultReq struct {

	// 所有任务信息
	TaskInfo []BatchConsistencyReq `json:"task_info"`
}

func (o BatchGetConsistencyResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetConsistencyResultReq struct{}"
	}

	return strings.Join([]string{"BatchGetConsistencyResultReq", string(data)}, " ")
}
