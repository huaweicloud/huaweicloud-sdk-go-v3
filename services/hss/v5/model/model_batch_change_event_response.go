package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeEventResponse Response Object
type BatchChangeEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchChangeEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeEventResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeEventResponse", string(data)}, " ")
}
