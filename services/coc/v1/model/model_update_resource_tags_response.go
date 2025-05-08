package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceTagsResponse Response Object
type UpdateResourceTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceTagsResponse", string(data)}, " ")
}
