package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesTagsByProjectIdResponse Response Object
type ListResourcesTagsByProjectIdResponse struct {

	// 标签列表
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListResourcesTagsByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesTagsByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesTagsByProjectIdResponse", string(data)}, " ")
}
