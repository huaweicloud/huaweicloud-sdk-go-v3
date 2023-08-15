package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchHandResponse Response Object
type BatchHandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchHandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchHandResponse struct{}"
	}

	return strings.Join([]string{"BatchHandResponse", string(data)}, " ")
}
