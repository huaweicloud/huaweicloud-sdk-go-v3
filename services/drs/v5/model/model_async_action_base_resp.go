package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 异步操作任务基础响应体。
type AsyncActionBaseResp struct {

	// 异步操作任务响应查询ID。
	QueryId string `json:"query_id"`
}

func (o AsyncActionBaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncActionBaseResp struct{}"
	}

	return strings.Join([]string{"AsyncActionBaseResp", string(data)}, " ")
}
