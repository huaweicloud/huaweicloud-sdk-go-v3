package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotIssueRequest 根据快照查询工作项信息的请求对象
type SnapshotIssueRequest struct {

	// 快照的ID数组。可以通过查询工作项快照列表接口获取，响应消息体中的id字段的值就是工作项快照ID。
	Ids []string `json:"ids"`

	// 是否返回工作项简要信息。 当值为false时ids中仅支持5个快照ID；值为true时，ids最多支持50个快照ID。
	SimpleResult *bool `json:"simple_result,omitempty"`
}

func (o SnapshotIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotIssueRequest struct{}"
	}

	return strings.Join([]string{"SnapshotIssueRequest", string(data)}, " ")
}
