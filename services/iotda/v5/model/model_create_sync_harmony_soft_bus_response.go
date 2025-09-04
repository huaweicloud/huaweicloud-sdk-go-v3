package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSyncHarmonySoftBusResponse Response Object
type CreateSyncHarmonySoftBusResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSyncHarmonySoftBusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSyncHarmonySoftBusResponse struct{}"
	}

	return strings.Join([]string{"CreateSyncHarmonySoftBusResponse", string(data)}, " ")
}
