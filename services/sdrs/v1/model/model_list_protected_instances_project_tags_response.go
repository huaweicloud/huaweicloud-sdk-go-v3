package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedInstancesProjectTagsResponse Response Object
type ListProtectedInstancesProjectTagsResponse struct {

	// 标签列表。
	Tags           *[]TagParams `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListProtectedInstancesProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesProjectTagsResponse", string(data)}, " ")
}
