package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagsResponse Response Object
type CreateResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceTagsResponse", string(data)}, " ")
}
