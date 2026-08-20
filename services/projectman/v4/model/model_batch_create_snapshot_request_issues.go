package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateSnapshotRequestIssues struct {

	// 工作项ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。 18~19个字符的数字字符串。
	Id string `json:"id"`
}

func (o BatchCreateSnapshotRequestIssues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSnapshotRequestIssues struct{}"
	}

	return strings.Join([]string{"BatchCreateSnapshotRequestIssues", string(data)}, " ")
}
