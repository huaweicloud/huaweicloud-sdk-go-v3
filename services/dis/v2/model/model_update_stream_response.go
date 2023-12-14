package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStreamResponse Response Object
type UpdateStreamResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamResponse struct{}"
	}

	return strings.Join([]string{"UpdateStreamResponse", string(data)}, " ")
}
