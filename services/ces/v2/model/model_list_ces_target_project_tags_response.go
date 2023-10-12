package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesTargetProjectTagsResponse Response Object
type ListCesTargetProjectTagsResponse struct {

	// 租户标签列表
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCesTargetProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesTargetProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCesTargetProjectTagsResponse", string(data)}, " ")
}
