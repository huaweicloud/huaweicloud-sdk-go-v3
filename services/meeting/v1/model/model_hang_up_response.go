package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type HangUpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HangUpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpResponse struct{}"
	}

	return strings.Join([]string{"HangUpResponse", string(data)}, " ")
}
