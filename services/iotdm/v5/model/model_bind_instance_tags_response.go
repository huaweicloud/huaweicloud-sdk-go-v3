package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindInstanceTagsResponse Response Object
type BindInstanceTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"BindInstanceTagsResponse", string(data)}, " ")
}
