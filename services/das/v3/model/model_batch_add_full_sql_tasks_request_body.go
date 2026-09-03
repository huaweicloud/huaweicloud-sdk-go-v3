package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddFullSqlTasksRequestBody 批量创建全量SQL明细解析任务请求体
type BatchAddFullSqlTasksRequestBody struct {

	// SQL解析任务列表
	QueryReqs []QueryReq `json:"query_reqs"`
}

func (o BatchAddFullSqlTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddFullSqlTasksRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddFullSqlTasksRequestBody", string(data)}, " ")
}
