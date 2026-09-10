package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeScheduledEventResponse Response Object
type AuthorizeScheduledEventResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AuthorizeScheduledEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeScheduledEventResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeScheduledEventResponse", string(data)}, " ")
}
