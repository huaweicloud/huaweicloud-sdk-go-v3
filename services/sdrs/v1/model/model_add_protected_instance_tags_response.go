package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddProtectedInstanceTagsResponse Response Object
type AddProtectedInstanceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddProtectedInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceTagsResponse", string(data)}, " ")
}
