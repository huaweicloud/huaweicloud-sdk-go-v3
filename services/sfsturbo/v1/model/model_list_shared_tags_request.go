package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSharedTagsRequest struct {
	// MIME类型

	ContentType string `json:"Content-Type"`
}

func (o ListSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedTagsRequest", string(data)}, " ")
}
