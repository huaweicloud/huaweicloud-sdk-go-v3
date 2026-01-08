package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeListenerTagsResponse Response Object
type ChangeListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"ChangeListenerTagsResponse", string(data)}, " ")
}
