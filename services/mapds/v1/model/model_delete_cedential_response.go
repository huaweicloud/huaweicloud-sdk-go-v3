package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCedentialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCedentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCedentialResponse struct{}"
	}

	return strings.Join([]string{"DeleteCedentialResponse", string(data)}, " ")
}
