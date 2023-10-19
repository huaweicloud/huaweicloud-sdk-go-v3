package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessClientResponse Response Object
type DeleteAccessClientResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAccessClientResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessClientResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccessClientResponse", string(data)}, " ")
}
