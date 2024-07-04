package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteQueuesResponse Response Object
type BatchDeleteQueuesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteQueuesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteQueuesResponse", string(data)}, " ")
}
