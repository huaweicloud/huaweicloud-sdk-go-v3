package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNodeResponse Response Object
type StopNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopNodeResponse struct{}"
	}

	return strings.Join([]string{"StopNodeResponse", string(data)}, " ")
}
