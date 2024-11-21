package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTagResponse Response Object
type BatchTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagResponse struct{}"
	}

	return strings.Join([]string{"BatchTagResponse", string(data)}, " ")
}
