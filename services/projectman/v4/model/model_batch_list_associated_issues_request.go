package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListAssociatedIssuesRequest Request Object
type BatchListAssociatedIssuesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o BatchListAssociatedIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListAssociatedIssuesRequest struct{}"
	}

	return strings.Join([]string{"BatchListAssociatedIssuesRequest", string(data)}, " ")
}
