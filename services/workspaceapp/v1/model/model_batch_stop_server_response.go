package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopServerResponse Response Object
type BatchStopServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchStopServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServerResponse struct{}"
	}

	return strings.Join([]string{"BatchStopServerResponse", string(data)}, " ")
}
