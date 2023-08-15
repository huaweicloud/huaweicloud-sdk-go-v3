package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePostalResponse Response Object
type DeletePostalResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePostalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePostalResponse struct{}"
	}

	return strings.Join([]string{"DeletePostalResponse", string(data)}, " ")
}
