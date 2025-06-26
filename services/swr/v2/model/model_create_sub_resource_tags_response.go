package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubResourceTagsResponse Response Object
type CreateSubResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateSubResourceTagsResponse", string(data)}, " ")
}
