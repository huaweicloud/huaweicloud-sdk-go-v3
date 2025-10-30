package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelectGroupAndStreamResponse Response Object
type SelectGroupAndStreamResponse struct {
	Job            *LtsInfoJob `json:"job,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o SelectGroupAndStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectGroupAndStreamResponse struct{}"
	}

	return strings.Join([]string{"SelectGroupAndStreamResponse", string(data)}, " ")
}
