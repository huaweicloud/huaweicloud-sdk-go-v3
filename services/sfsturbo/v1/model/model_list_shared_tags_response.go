package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSharedTagsResponse struct {
	// tag标签的列表

	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSharedTagsResponse", string(data)}, " ")
}
