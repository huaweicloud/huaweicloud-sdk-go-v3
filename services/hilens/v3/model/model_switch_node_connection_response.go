package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchNodeConnectionResponse Response Object
type SwitchNodeConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchNodeConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchNodeConnectionResponse struct{}"
	}

	return strings.Join([]string{"SwitchNodeConnectionResponse", string(data)}, " ")
}
