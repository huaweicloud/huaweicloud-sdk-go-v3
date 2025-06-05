package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartProtectionResponse Response Object
type BatchStartProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchStartProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartProtectionResponse struct{}"
	}

	return strings.Join([]string{"BatchStartProtectionResponse", string(data)}, " ")
}
