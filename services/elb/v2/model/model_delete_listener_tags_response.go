package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsResponse", string(data)}, " ")
}
