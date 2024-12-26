package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTagsResponse Response Object
type SaveTagsResponse struct {
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTagsResponse struct{}"
	}

	return strings.Join([]string{"SaveTagsResponse", string(data)}, " ")
}
