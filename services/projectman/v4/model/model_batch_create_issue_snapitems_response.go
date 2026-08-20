package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIssueSnapitemsResponse Response Object
type BatchCreateIssueSnapitemsResponse struct {
	Result *BatchCreateSnapshotResponseResult `json:"result,omitempty"`

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 失败原因。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateIssueSnapitemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIssueSnapitemsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateIssueSnapitemsResponse", string(data)}, " ")
}
