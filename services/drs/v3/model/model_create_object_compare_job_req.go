package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectCompareJobReq 创建对象级对比任务请求。
type CreateObjectCompareJobReq struct {

	// 对比任务线程数量，当前仅cloudDataGuard-cassandra和cloudDataGuard-gausscassandra-to-gausscassandra链路支持。
	CompareTaskNum *int32 `json:"compare_task_num,omitempty"`
}

func (o CreateObjectCompareJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectCompareJobReq struct{}"
	}

	return strings.Join([]string{"CreateObjectCompareJobReq", string(data)}, " ")
}
