package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeletePropertyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePropertyResponse struct{}"
	}

	return strings.Join([]string{"DeletePropertyResponse", string(data)}, " ")
}
