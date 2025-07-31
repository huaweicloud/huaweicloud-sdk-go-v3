package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartWebTamperProtectionResponse Response Object
type BatchStartWebTamperProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchStartWebTamperProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartWebTamperProtectionResponse struct{}"
	}

	return strings.Join([]string{"BatchStartWebTamperProtectionResponse", string(data)}, " ")
}
