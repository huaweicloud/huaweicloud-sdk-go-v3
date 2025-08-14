package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSessionsResponse Response Object
type BatchDeleteSessionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSessionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSessionsResponse", string(data)}, " ")
}
