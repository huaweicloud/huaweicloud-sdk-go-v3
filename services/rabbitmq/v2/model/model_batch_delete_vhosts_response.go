package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteVhostsResponse Response Object
type BatchDeleteVhostsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVhostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVhostsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVhostsResponse", string(data)}, " ")
}
