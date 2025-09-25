package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMessageAdditionsResponse Response Object
type CreateMessageAdditionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMessageAdditionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageAdditionsResponse struct{}"
	}

	return strings.Join([]string{"CreateMessageAdditionsResponse", string(data)}, " ")
}
