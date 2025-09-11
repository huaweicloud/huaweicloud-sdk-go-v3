package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceProjectTagsResponse Response Object
type ListInstanceProjectTagsResponse struct {

	// 标签列表
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListInstanceProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceProjectTagsResponse", string(data)}, " ")
}
