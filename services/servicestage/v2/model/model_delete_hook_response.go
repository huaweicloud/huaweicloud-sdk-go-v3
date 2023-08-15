package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHookResponse Response Object
type DeleteHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHookResponse struct{}"
	}

	return strings.Join([]string{"DeleteHookResponse", string(data)}, " ")
}
