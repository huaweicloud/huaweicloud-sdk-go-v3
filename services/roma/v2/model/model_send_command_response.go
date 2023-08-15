package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendCommandResponse Response Object
type SendCommandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendCommandResponse struct{}"
	}

	return strings.Join([]string{"SendCommandResponse", string(data)}, " ")
}
