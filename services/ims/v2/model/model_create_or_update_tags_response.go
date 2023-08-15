package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrUpdateTagsResponse Response Object
type CreateOrUpdateTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateOrUpdateTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateTagsResponse", string(data)}, " ")
}
