package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStackResponse Response Object
type DeleteStackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackResponse struct{}"
	}

	return strings.Join([]string{"DeleteStackResponse", string(data)}, " ")
}
