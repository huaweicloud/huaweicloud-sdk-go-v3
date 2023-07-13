package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLayoutResponse Response Object
type DeleteLayoutResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLayoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLayoutResponse struct{}"
	}

	return strings.Join([]string{"DeleteLayoutResponse", string(data)}, " ")
}
