package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCommandResponse Response Object
type DeleteCommandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCommandResponse struct{}"
	}

	return strings.Join([]string{"DeleteCommandResponse", string(data)}, " ")
}
