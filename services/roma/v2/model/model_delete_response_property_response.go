package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResponsePropertyResponse Response Object
type DeleteResponsePropertyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResponsePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResponsePropertyResponse struct{}"
	}

	return strings.Join([]string{"DeleteResponsePropertyResponse", string(data)}, " ")
}
