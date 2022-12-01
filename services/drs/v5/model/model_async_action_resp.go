package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 异步操作任务响应体。
type AsyncActionResp struct {

	// 异步操作任务响应查询ID。
	QueryId string `json:"query_id"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`
}

func (o AsyncActionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncActionResp struct{}"
	}

	return strings.Join([]string{"AsyncActionResp", string(data)}, " ")
}
