package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectProtectedTagsResponse Response Object
type ListProjectProtectedTagsResponse struct {

	// 保护tag详情列表
	Body           *[]ProjectProtectedTagPo `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListProjectProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectProtectedTagsResponse", string(data)}, " ")
}
